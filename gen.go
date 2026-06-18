//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
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
	Type        string                 `json:"type"`
	Properties  map[string]interface{} `json:"properties"`
	Description string                 `json:"description"`
	Title       string                 `json:"title"`
}

type PathItem struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}
type ReqBody struct {
	Content map[string]struct {
		Schema struct{ Ref string `json:"$ref"` } `json:"schema"`
	} `json:"content"`
}

func containsAny(s string, list []string) bool {
	for _, item := range list {
		if s == item { return true }
	}
	return false
}

func truncateRunes(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen]) + "..."
}

func sanitizeComment(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	var clean []byte
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size <= 1 {
			i++
			continue
		}
		clean = append(clean, []byte(string(r))...)
		i += size
	}
	s = string(clean)
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return strings.TrimSpace(truncateRunes(s, 120))
}

func appendType(out *[]string, name string, s Schema, cache map[string]string) {
	if s.Description != "" {
		*out = append(*out, fmt.Sprintf("// %s", sanitizeComment(s.Description)))
	}
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
		var p struct{ Ref, Type, Format, Description string }
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
		line := fmt.Sprintf("\t%s %s `json:\"%s\"`", fn, ft, jtag)
		if p.Description != "" {
			comment := sanitizeComment(p.Description)
			if comment != "" {
				line += fmt.Sprintf(" // %s", comment)
			}
		}
		*out = append(*out, line)
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

	depOf := map[string]map[string]bool{}
	for sname, s := range schemas {
		depOf[sname] = map[string]bool{}
		if s.Properties == nil { continue }
		for _, pv := range s.Properties {
			pr, _ := json.Marshal(pv)
			var p struct{ Ref string }
			json.Unmarshal(pr, &p)
			if p.Ref != "" {
				rn := strings.TrimPrefix(p.Ref, "#/components/schemas/")
				if rn != sname { depOf[sname][rn] = true }
			}
		}
	}

	type Method struct {
		Dir, Name, HTTPM, Path, ReqT, RespT, Summary string
		Deprecated  bool
		ReplaceWith string
		Notes       []string
	}
	dirMethods := map[string][]Method{}
	dirDirectTypes := map[string]map[string]bool{}
	methodNames := map[string]bool{}

	for path, items := range paths {
		for httpM, raw := range items {
			var item struct {
				Summary     string                    `json:"summary"`
				Description string                    `json:"description"`
				OperationID string                    `json:"operationId"`
				Tags        []string                  `json:"tags"`
				Deprecated  bool                      `json:"deprecated"`
				RequestBody *json.RawMessage           `json:"requestBody"`
				Responses   map[string]json.RawMessage `json:"responses"`
			}
			json.Unmarshal(raw, &item)
			dir := ""
			for _, t := range item.Tags {
				if d, ok := tagDir[t]; ok { dir = d; break }
			}
			if dir == "" { continue }
			if dirDirectTypes[dir] == nil { dirDirectTypes[dir] = map[string]bool{} }

			reqT := ""
			if item.RequestBody != nil {
				var rb ReqBody
				json.Unmarshal(*item.RequestBody, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					sname := strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")
					reqT = cache[sname]
					dirDirectTypes[dir][sname] = true
				}
			}
			respT := ""
			if r, ok := item.Responses["200"]; ok {
				var rb ReqBody
				json.Unmarshal(r, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					sname := strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")
					respT = cache[sname]
					dirDirectTypes[dir][sname] = true
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
			replaceWith := ""
			desc := item.Description
			isDep := item.Deprecated || strings.Contains(desc, "停用") || strings.Contains(desc, "弃用")
			if isDep && !item.Deprecated {
				item.Deprecated = true
			}
			if isDep {
				if idx := strings.Index(desc, "\">/"); idx > 0 {
					pathStart := idx + 2
					pathEnd := strings.Index(desc[pathStart:], "</a>")
					if pathEnd > 0 {
						replaceWith = desc[pathStart : pathStart+pathEnd]
					}
				}
			}

			// Extract important usage notes from description
			var notes []string
			noteKeywords := []string{"请使用", "请确保", "请注意", "每30秒", "每", "状态码为200",
				"不能", "无法", "不可", "必须先", "请先",
				"每", "如需", "要获取", "为了", "在更新之前", "在更改状态前"}
			stripRE := regexp.MustCompile(`<[^>]+>`)
			for _, kw := range noteKeywords {
				idx := strings.Index(desc, kw)
				if idx < 0 { continue }
				start := idx
				end := idx + 130
				if end > len(desc) { end = len(desc) }
				note := desc[start:end]
				if nl := strings.IndexAny(note, ".\n。"); nl > 20 {
					note = note[:nl+1]
				}
				note = stripRE.ReplaceAllString(note, "")
				note = strings.TrimSpace(note)
				note = sanitizeComment(note)
				if len(note) > 15 && !containsAny(note, notes) &&
					!strings.Contains(note, "停用") && !strings.Contains(note, "弃用") {
					notes = append(notes, note)
				}
			}
			// Deduplicate and limit
			if len(notes) > 3 {
				notes = notes[:3]
			}

			dirMethods[dir] = append(dirMethods[dir], Method{
				Dir: dir, Name: name, HTTPM: strings.ToUpper(httpM),
				Path: path, ReqT: reqT, RespT: respT,
				Summary: item.Summary, Deprecated: item.Deprecated,
				ReplaceWith: replaceWith, Notes: notes,
			})
		}
	}

	dirAllTypes := map[string]map[string]bool{}
	for dir, direct := range dirDirectTypes {
		visited := map[string]bool{}
		var visit func(string)
		visit = func(sname string) {
			if visited[sname] { return }
			visited[sname] = true
			for dep := range depOf[sname] { visit(dep) }
		}
		for s := range direct { visit(s) }
		dirAllTypes[dir] = visited
	}

	typeDirCount := map[string]int{}
	for _, types := range dirAllTypes {
		for t := range types { typeDirCount[t]++ }
	}

	sharedTypes := map[string]bool{}
	for t, c := range typeDirCount {
		if c > 1 { sharedTypes[t] = true }
	}

	forcedShared := map[string]bool{"rpcStatus": true, "googlerpcStatus": true, "protobufAny": true}
	for t := range forcedShared { sharedTypes[t] = true }

	os.MkdirAll("types", 0755)
	genned := map[string]bool{}
	inProg := map[string]bool{}
	nameUsed := map[string]string{}
	tlines := []string{"package types", ""}

	var genType func(string, *[]string)
	genType = func(sname string, out *[]string) {
		if genned[sname] || inProg[sname] { return }
		inProg[sname] = true
		s, ok := schemas[sname]
		if !ok { inProg[sname] = false; return }
		gn := cache[sname]
		if prev, exists := nameUsed[gn]; exists && prev != sname {
			inProg[sname] = false; genned[sname] = true; return
		}
		nameUsed[gn] = sname
		for dep := range depOf[sname] {
			if dep != sname { genType(dep, out) }
		}
		inProg[sname] = false
		genned[sname] = true
		appendType(out, gn, s, cache)
	}

	for sname := range sharedTypes { genType(sname, &tlines) }
	os.WriteFile("types/types.go", []byte(strings.Join(tlines, "\n")), 0644)
	fmt.Printf("types/types.go: %d shared types\n", len(genned))

	for dir, methods := range dirMethods {
		dp := filepath.Join(dir)
		os.MkdirAll(dp, 0755)

		dlines := []string{fmt.Sprintf("package %s", dir), ""}
		typeCount := 0
		for t := range dirAllTypes[dir] {
			if sharedTypes[t] || genned[t] { continue }
			if _, ok := schemas[t]; !ok { continue }
			var writeWithDeps func(string)
			writeWithDeps = func(sname string) {
				if genned[sname] || sharedTypes[sname] { return }
				s2, ok2 := schemas[sname]
				if !ok2 { return }
				for dep := range depOf[sname] {
					if !sharedTypes[dep] && !genned[dep] { writeWithDeps(dep) }
				}
				appendType(&dlines, cache[sname], s2, cache)
				genned[sname] = true
			}
			writeWithDeps(t)
			typeCount++
		}

		if typeCount > 0 {
			os.WriteFile(dp+"/types.go", []byte(strings.Join(dlines, "\n")), 0644)
		}

		needsTypes := false
		for _, m := range methods {
			for _, t := range []string{m.ReqT, m.RespT} {
				if t == "" { continue }
				for sname := range sharedTypes {
					if cache[sname] == t { needsTypes = true; break }
				}
			}
		}

		var ml []string
		ml = append(ml, fmt.Sprintf("package %s", dir), "")
		imports := `"context"; "github.com/QuoVadis86/go-ozon-sdk/transport"`
		if needsTypes { imports += `; "github.com/QuoVadis86/go-ozon-sdk/types"` }
		ml = append(ml, fmt.Sprintf("import (%s)", imports), "")
		ml = append(ml, "type Service struct { Client *transport.Client }", "")
		for _, m := range methods {
			if m.Deprecated {
				depMsg := "use " + m.ReplaceWith + " instead"
				if m.ReplaceWith == "" {
					depMsg = "this method is deprecated and will be removed"
				}
				ml = append(ml, fmt.Sprintf("// Deprecated: %s", sanitizeComment(depMsg)))
			}
			if m.Summary != "" {
				ml = append(ml, fmt.Sprintf("// %s", sanitizeComment(m.Summary)))
			}
			for _, note := range m.Notes {
				ml = append(ml, fmt.Sprintf("// Note: %s", note))
			}
			fn := fmt.Sprintf("func (s *Service) %s(ctx context.Context", m.Name)
			pkg := ""
			if m.ReqT != "" {
				isShared := false
				for sname := range sharedTypes {
					if cache[sname] == m.ReqT { isShared = true; break }
				}
				if isShared { pkg = "types." }
				fn += fmt.Sprintf(", req *%s%s", pkg, m.ReqT)
			}
			fn += ") "
			if m.RespT != "" {
				isShared := false
				for sname := range sharedTypes {
					if cache[sname] == m.RespT { isShared = true; break }
				}
				pkg2 := ""
				if isShared { pkg2 = "types." }
				fn += fmt.Sprintf("(*%s%s, error)", pkg2, m.RespT)
			} else { fn += "error" }
			ml = append(ml, fn+" {")
			if m.RespT != "" {
				isShared := false
				for sname := range sharedTypes {
					if cache[sname] == m.RespT { isShared = true; break }
				}
				pkg2 := ""
				if isShared { pkg2 = "types." }
				ml = append(ml, fmt.Sprintf("\tvar resp %s%s", pkg2, m.RespT))
			}
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

		// Generate test stub
		tlines := []string{
			fmt.Sprintf("package %s", dir), "",
			`import "testing"`, "",
			"func TestService_New(t *testing.T) {",
			"\tsvc := &Service{Client: nil}",
			"\t_ = svc",
			"}",
			"",
		}
		os.WriteFile(dp+"/service_test.go", []byte(strings.Join(tlines, "\n")), 0644)
		fmt.Printf("%s: %d methods, %d types\n", dir, len(methods), typeCount)
	}
}
