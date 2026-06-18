//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	camelRe = regexp.MustCompile(`[._\-/]`)
	abrMap  = map[string]string{
		"id": "ID", "sku": "SKU", "url": "URL", "json": "JSON",
		"html": "HTML", "api": "API", "fbs": "FBS", "fbo": "FBO",
		"rfbs": "RFBS", "pdf": "PDF", "png": "PNG", "ettn": "ETTN",
		"gtd": "GTD", "ttn": "TTN", "tpl": "TPL", "dlv": "DLV",
		"erfbs": "ERFBS", "fbp": "FBP",
	}
	tagDir = map[string]string{
		"SellerInfo": "seller", "APIkey": "seller",
		"CategoryAPI": "category", "ProductAPI": "product",
		"Prices&StocksAPI": "prices", "BarcodeAPI": "barcode",
		"WarehouseAPI": "warehouse", "FBSWarehouseSetup": "warehouse",
		"FBS": "fbs", "FBS&rFBSMarks": "fbs", "DeliveryFBS": "fbs", "DeliveryrFBS": "fbs",
		"FBO": "fbo", "DeliveryFBP": "fbo", "DeliveryFBPDraft": "fbo",
		"DraftDirectFBP": "fbo", "DraftDropOffFBP": "fbo", "DraftPickupFBP": "fbo",
		"OrderDirectFBP": "fbo", "OrderDropOffFBP": "fbo", "OrderPickupFBP": "fbo",
		"FinanceAPI": "finance", "ReportAPI": "report",
		"Promos": "promo", "PromosBeta": "promo", "SellerActions": "promo",
		"PricingStrategyAPI": "pricing",
		"ReviewAPI": "review", "Questions&Answers": "review",
		"ChatAPI": "chat", "Notification": "notification",
		"ReturnsAPI": "returns", "RFBSReturnsAPI": "returns",
		"ReturnAPI": "returns", "CancellationAPI": "returns",
		"Pass": "pass", "SellerRating": "rating", "PolygonAPI": "delivery",
		"BetaMethod": "beta", "Examples": "beta",
		"Premium": "premium", "BrandAPI": "certificate", "CertificationAPI": "certificate",
	}
)

func toCamel(s string) string {
	parts := camelRe.Split(s, -1)
	var r string
	for _, p := range parts {
		if p == "" { continue }
		l := strings.ToLower(p)
		if v, ok := abrMap[l]; ok { r += v } else { r += strings.ToUpper(p[:1]) + p[1:] }
	}
	if r == "" { return "X" }
	return r
}

func toGoName(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	parts := strings.Split(s, "_")
	var r string
	for _, p := range parts {
		if p == "" { continue }
		l := strings.ToLower(p)
		if v, ok := abrMap[l]; ok { r += v } else { r += strings.ToUpper(p[:1]) + p[1:] }
	}
	if r != "" && r[0] >= '0' && r[0] <= '9' { r = "X" + r }
	return r
}

type Schema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}

type ReqBody struct {
	Content map[string]struct {
		Schema struct{ Ref string `json:"$ref"` } `json:"schema"`
	} `json:"content"`
}

func appendType(out *[]string, name string, s Schema, cache map[string]string) {
	if s.Properties == nil || len(s.Properties) == 0 {
		switch s.Type {
		case "array": *out = append(*out, fmt.Sprintf("type %s []interface{}", name))
		case "string": *out = append(*out, fmt.Sprintf("type %s string", name))
		case "integer": *out = append(*out, fmt.Sprintf("type %s int64", name))
		case "number": *out = append(*out, fmt.Sprintf("type %s float64", name))
		case "boolean": *out = append(*out, fmt.Sprintf("type %s bool", name))
		default: *out = append(*out, fmt.Sprintf("type %s interface{}", name))
		}
		*out = append(*out, "")
		return
	}
	*out = append(*out, fmt.Sprintf("type %s struct {", name))
	for pn, pv := range s.Properties {
		pr, _ := json.Marshal(pv)
		var p struct{ Ref, Type, Format string }
		json.Unmarshal(pr, &p)
		fn := toGoName(pn)
		ft := "interface{}"
		if p.Ref != "" { ft = cache[strings.TrimPrefix(p.Ref, "#/components/schemas/")]
		} else if p.Type == "array" { ft = "[]interface{}"
		} else if p.Type == "string" { ft = "string"
		} else if p.Type == "integer" {
			if p.Format == "int32" { ft = "int32" } else { ft = "int64" }
		} else if p.Type == "number" { ft = "float64"
		} else if p.Type == "boolean" { ft = "bool"
		} else if p.Type == "object" { ft = "map[string]interface{}" }
		jtag := pn
		if pn == "type" { jtag = "type_" }
		*out = append(*out, fmt.Sprintf("\t%s %s `json:\"%s\"`", fn, ft, jtag))
	}
	*out = append(*out, "}", "")
}

func main() {
	data, err := os.ReadFile("origin/swagger.json")
	if err != nil { panic(err) }
	var raw map[string]json.RawMessage
	json.Unmarshal(data, &raw)
	var paths map[string]map[string]json.RawMessage
	json.Unmarshal(raw["paths"], &paths)
	var comps struct{ Schemas map[string]json.RawMessage `json:"schemas"` }
	json.Unmarshal(raw["components"], &comps)

	schemas := map[string]Schema{}
	for n, raw := range comps.Schemas {
		var s Schema; json.Unmarshal(raw, &s); schemas[n] = s
	}
	cache := map[string]string{}
	for n := range comps.Schemas { cache[n] = toCamel(n) }

	type Method struct{ Dir, Name, HTTPM, Path, ReqT, RespT string }
	dirMethods := map[string][]Method{}
	methodNames := map[string]bool{}

	for path, items := range paths {
		for httpM, raw := range items {
			var item struct {
				OperationID string                    `json:"operationId"`
				Tags        []string                  `json:"tags"`
				RequestBody *json.RawMessage           `json:"requestBody"`
				Responses   map[string]json.RawMessage `json:"responses"`
			}
			json.Unmarshal(raw, &item)
			dir := ""
			for _, t := range item.Tags {
				if d, ok := tagDir[t]; ok { dir = d; break }
			}
			if dir == "" { continue }

			reqT := ""
			if item.RequestBody != nil {
				var rb ReqBody
				json.Unmarshal(*item.RequestBody, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					reqT = cache[strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")]
				}
			}
			respT := ""
			if r, ok := item.Responses["200"]; ok {
				var rb ReqBody
				json.Unmarshal(r, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					respT = cache[strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")]
				}
			}

			name := item.OperationID
			parts := strings.Split(name, "_")
			if len(parts) > 1 { name = toGoName(strings.Join(parts[1:], "_"))
			} else { name = toGoName(name) }

			k := dir + ":" + name
			if _, exists := methodNames[k]; exists {
				for _, p := range strings.Split(strings.Trim(path, "/"), "/") {
					if !strings.ContainsAny(p, "{}") { name += toGoName(p); break }
				}
				if _, exists := methodNames[dir+":"+name]; exists {
					name += "V" + strings.ReplaceAll(path, "/", "_")
				}
			}
			methodNames[dir+":"+name] = true
			dirMethods[dir] = append(dirMethods[dir], Method{dir, name, strings.ToUpper(httpM), path, reqT, respT})
		}
	}

	// 1. Generate types/types.go
	os.MkdirAll("types", 0755)
	genned, inProg := map[string]bool{}, map[string]bool{}
	nameUsed := map[string]string{}
	tlines := []string{"package types", ""}

	var genType func(string)
	genType = func(sname string) {
		if genned[sname] || inProg[sname] { return }
		inProg[sname] = true
		s, ok := schemas[sname]
		if !ok { inProg[sname] = false; return }
		gn := cache[sname]
		if prev, exists := nameUsed[gn]; exists && prev != sname {
			inProg[sname] = false; genned[sname] = true; return
		}
		nameUsed[gn] = sname
		if s.Properties != nil {
			for _, pv := range s.Properties {
				pr, _ := json.Marshal(pv)
				var p struct{ Ref string }
				json.Unmarshal(pr, &p)
				if p.Ref != "" {
					rn := strings.TrimPrefix(p.Ref, "#/components/schemas/")
					if rn != sname { genType(rn) }
				}
			}
		}
		inProg[sname] = false
		genned[sname] = true
		appendType(&tlines, gn, s, cache)
	}
	for sname := range schemas { genType(sname) }
	os.WriteFile("types/types.go", []byte(strings.Join(tlines, "\n")), 0644)
	fmt.Printf("types/types.go: %d types\n", len(genned))

	// 2. Generate per-service service.go at module root
	for dir, methods := range dirMethods {
		dp := filepath.Join(dir)
		os.MkdirAll(dp, 0755)

		var ml []string
		ml = append(ml, fmt.Sprintf("package %s", dir), "")
		ml = append(ml, `import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")`, "")
		ml = append(ml, "type Service struct { Client *internal.Client }", "")
		for _, m := range methods {
			pkg := "types."
			fn := fmt.Sprintf("func (s *Service) %s(ctx context.Context", m.Name)
			if m.ReqT != "" { fn += fmt.Sprintf(", req *%s%s", pkg, m.ReqT) }
			fn += ") "
			if m.RespT != "" { fn += fmt.Sprintf("(*%s%s, error)", pkg, m.RespT) } else { fn += "error" }
			ml = append(ml, fn+" {")
			if m.RespT != "" { ml = append(ml, fmt.Sprintf("\tvar resp %s%s", pkg, m.RespT)) }
			meth := "Post"
			if m.HTTPM == "GET" { meth = "Get" }
			call := fmt.Sprintf("\terr := s.Client.%s(ctx, %q", meth, m.Path)
			if m.ReqT != "" {
				if m.RespT != "" { call += ", req, &resp)" } else { call += ", req, nil)" }
			} else {
				if m.RespT != "" { call += ", nil, &resp)" } else { call += ", nil, nil)" }
			}
			ml = append(ml, call)
			ml = append(ml, "\tif err != nil {")
			if m.RespT != "" { ml = append(ml, "\t\treturn nil, err") } else { ml = append(ml, "\t\treturn err") }
			ml = append(ml, "\t}")
			if m.RespT != "" { ml = append(ml, "\treturn &resp, nil") } else { ml = append(ml, "\treturn nil") }
			ml = append(ml, "}", "")
		}
		os.WriteFile(dp+"/service.go", []byte(strings.Join(ml, "\n")), 0644)
		fmt.Printf("%s: %d methods\n", dir, len(methods))
	}
}
