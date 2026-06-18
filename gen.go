//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var camelRe = regexp.MustCompile(`[._\-/]`)

var abrMap = map[string]string{
	"id": "ID", "sku": "SKU", "url": "URL", "json": "JSON",
	"html": "HTML", "api": "API", "fbs": "FBS", "fbo": "FBO",
	"rfbs": "RFBS", "pdf": "PDF", "png": "PNG", "ettn": "ETTN",
	"gtd": "GTD", "ttn": "TTN", "tpl": "TPL", "dlv": "DLV",
	"erfbs": "ERFBS", "fbp": "FBP",
}

func toCamel(s string) string {
	parts := camelRe.Split(s, -1)
	var r string
	for _, p := range parts {
		if p == "" {
			continue
		}
		l := strings.ToLower(p)
		if v, ok := abrMap[l]; ok {
			r += v
		} else {
			r += strings.ToUpper(p[:1]) + p[1:]
		}
	}
	if r == "" {
		return "X"
	}
	return r
}

func toGoName(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	parts := strings.Split(s, "_")
	var r string
	for _, p := range parts {
		if p == "" {
			continue
		}
		l := strings.ToLower(p)
		if v, ok := abrMap[l]; ok {
			r += v
		} else {
			r += strings.ToUpper(p[:1]) + p[1:]
		}
	}
	if r != "" && r[0] >= '0' && r[0] <= '9' {
		r = "X" + r
	}
	return r
}

var tagSvc = map[string]string{
	"SellerInfo": "Seller", "APIkey": "Seller",
	"CategoryAPI": "Category",
	"ProductAPI": "Product",
	"Prices&StocksAPI": "Prices",
	"BarcodeAPI": "Barcode",
	"WarehouseAPI": "Warehouse", "FBSWarehouseSetup": "Warehouse",
	"FBS": "FBS", "FBS&rFBSMarks": "FBS", "DeliveryFBS": "FBS", "DeliveryrFBS": "FBS",
	"FBO": "FBO", "DeliveryFBP": "FBO", "DeliveryFBPDraft": "FBO",
	"DraftDirectFBP": "FBO", "DraftDropOffFBP": "FBO", "DraftPickupFBP": "FBO",
	"OrderDirectFBP": "FBO", "OrderDropOffFBP": "FBO", "OrderPickupFBP": "FBO",
	"FinanceAPI": "Finance", "ReportAPI": "Report",
	"Promos": "Promo", "PromosBeta": "Promo", "SellerActions": "Promo",
	"PricingStrategyAPI": "Pricing",
	"ReviewAPI": "Review", "Questions&Answers": "Review",
	"ChatAPI": "Chat", "Notification": "Notification",
	"ReturnsAPI": "Returns", "RFBSReturnsAPI": "Returns",
	"ReturnAPI": "Returns", "CancellationAPI": "Returns",
	"Pass": "Pass", "SellerRating": "Rating", "PolygonAPI": "Delivery",
	"BetaMethod": "Beta", "Examples": "Beta",
	"Premium": "Premium", "BrandAPI": "Certificate", "CertificationAPI": "Certificate",
}

type Schema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Items      interface{}            `json:"items"`
	Format     string                 `json:"format"`
}

type ReqBody struct {
	Content map[string]struct {
		Schema struct{ Ref string `json:"$ref"` } `json:"schema"`
	} `json:"content"`
}

func main() {
	data, err := os.ReadFile("origin/swagger.json")
	if err != nil {
		panic(err)
	}

	var rawSpec map[string]json.RawMessage
	json.Unmarshal(data, &rawSpec)

	var paths map[string]map[string]json.RawMessage
	json.Unmarshal(rawSpec["paths"], &paths)

	var comps struct{ Schemas map[string]json.RawMessage `json:"schemas"` }
	json.Unmarshal(rawSpec["components"], &comps)

	schemaMap := map[string]Schema{}
	for n, raw := range comps.Schemas {
		var s Schema
		json.Unmarshal(raw, &s)
		schemaMap[n] = s
	}

	typeNameCache := map[string]string{}
	for n := range comps.Schemas {
		typeNameCache[n] = toCamel(n)
	}

	type Method struct {
		Service, Name, HTTPMethod, Path, RequestType, ResponseType string
	}

	methods := map[string][]Method{}
	methodNames := map[string]bool{}

	for path, methodsRaw := range paths {
		for httpMethod, itemRaw := range methodsRaw {
			var item struct {
				Summary     string                    `json:"summary"`
				OperationID string                    `json:"operationId"`
				Tags        []string                  `json:"tags"`
				RequestBody *json.RawMessage           `json:"requestBody"`
				Responses   map[string]json.RawMessage `json:"responses"`
			}
			json.Unmarshal(itemRaw, &item)

			tag := ""
			for _, t := range item.Tags {
				if _, ok := tagSvc[t]; ok {
					tag = t
					break
				}
			}
			if tag == "" {
				continue
			}
			svc := tagSvc[tag]

			reqType := ""
			if item.RequestBody != nil {
				var rb ReqBody
				json.Unmarshal(*item.RequestBody, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					reqType = toCamel(strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/"))
				}
			}
			respType := ""
			if resp, ok := item.Responses["200"]; ok {
				var rb ReqBody
				json.Unmarshal(resp, &rb)
				if c, ok := rb.Content["application/json"]; ok && c.Schema.Ref != "" {
					respType = toCamel(strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/"))
				}
			}

			name := item.OperationID
			parts := strings.Split(name, "_")
			if len(parts) > 1 {
				name = toGoName(strings.Join(parts[1:], "_"))
			} else {
				name = toGoName(name)
			}

			key := svc + ":" + name
			if _, exists := methodNames[key]; exists {
				pathParts := strings.Split(strings.Trim(path, "/"), "/")
				suffix := ""
				for _, pp := range pathParts {
					if !strings.ContainsAny(pp, "{}") {
						suffix = toGoName(pp)
					}
				}
				if suffix != "" {
					name += suffix
				}
				key2 := svc + ":" + name
				if _, exists := methodNames[key2]; exists {
					name += "V" + strings.ReplaceAll(path, "/", "_")
				}
			}
			methodNames[svc+":"+name] = true

			methods[svc] = append(methods[svc], Method{
				Service: svc, Name: name, HTTPMethod: strings.ToUpper(httpMethod),
				Path: path, RequestType: reqType, ResponseType: respType,
			})
		}
	}

	typeLines := []string{"package ozon", ""}
	genned := map[string]bool{}
	generating := map[string]bool{}
	goNameUsed := map[string]string{}

	var writeType func(sname string)
	writeType = func(sname string) {
		if genned[sname] {
			return
		}
		if generating[sname] {
			return
		}
		generating[sname] = true
		genned[sname] = true

		schema, ok := schemaMap[sname]
		if !ok {
			generating[sname] = false
			return
		}
		goName := typeNameCache[sname]

		if _, exists := goNameUsed[goName]; exists && goNameUsed[goName] != sname {
			generating[sname] = false
			return
		}
		goNameUsed[goName] = sname

		if schema.Properties != nil {
			for _, pval := range schema.Properties {
				propRaw, _ := json.Marshal(pval)
				var prop struct {
					Ref  string `json:"$ref"`
					Type string `json:"type"`
				}
				json.Unmarshal(propRaw, &prop)
				if prop.Ref != "" {
					refName := strings.TrimPrefix(prop.Ref, "#/components/schemas/")
					if refName != sname {
						writeType(refName)
					}
				}
			}
		}
		generating[sname] = false

		if schema.Properties == nil || len(schema.Properties) == 0 {
			t := schema.Type
			switch t {
			case "array":
				typeLines = append(typeLines, fmt.Sprintf("type %s []interface{}", goName))
			case "string":
				typeLines = append(typeLines, fmt.Sprintf("type %s string", goName))
			case "integer":
				typeLines = append(typeLines, fmt.Sprintf("type %s int64", goName))
			case "number":
				typeLines = append(typeLines, fmt.Sprintf("type %s float64", goName))
			case "boolean":
				typeLines = append(typeLines, fmt.Sprintf("type %s bool", goName))
			default:
				typeLines = append(typeLines, fmt.Sprintf("type %s interface{}", goName))
			}
			typeLines = append(typeLines, "")
			return
		}

		typeLines = append(typeLines, fmt.Sprintf("type %s struct {", goName))
		for pname, pval := range schema.Properties {
			propRaw, _ := json.Marshal(pval)
			var prop struct {
				Ref    string `json:"$ref"`
				Type   string `json:"type"`
				Format string `json:"format"`
			}
			json.Unmarshal(propRaw, &prop)

			fieldName := toGoName(pname)
			fieldType := "interface{}"
			if prop.Ref != "" {
				refName := strings.TrimPrefix(prop.Ref, "#/components/schemas/")
				fieldType = typeNameCache[refName]
			} else if prop.Type == "array" {
				fieldType = "[]interface{}"
			} else if prop.Type == "string" {
				fieldType = "string"
			} else if prop.Type == "integer" {
				if prop.Format == "int32" {
					fieldType = "int32"
				} else {
					fieldType = "int64"
				}
			} else if prop.Type == "number" {
				fieldType = "float64"
			} else if prop.Type == "boolean" {
				fieldType = "bool"
			} else if prop.Type == "object" {
				fieldType = "map[string]interface{}"
			}

			jsonTag := pname
			if pname == "type" {
				jsonTag = "type_"
			}
			typeLines = append(typeLines, fmt.Sprintf("\t%s %s `json:\"%s\"`", fieldName, fieldType, jsonTag))
		}
		typeLines = append(typeLines, "}", "")
	}

	for sname := range schemaMap {
		writeType(sname)
	}

	os.WriteFile("ozon/types.go", []byte(strings.Join(typeLines, "\n")), 0644)
	fmt.Printf("ozon/types.go: %d types\n", len(genned))

	for svc, mlist := range methods {
		implFile := fmt.Sprintf("ozon/%s_impl.go", strings.ToLower(svc))
		var lines []string
		lines = append(lines, "package ozon", "", `import "context"`, "")
		for _, m := range mlist {
			fnSig := fmt.Sprintf("func (s *%sService) %s(ctx context.Context", svc, m.Name)
			if m.RequestType != "" {
				fnSig += fmt.Sprintf(", req *%s", m.RequestType)
			}
			fnSig += ") "
			if m.ResponseType != "" {
				fnSig += fmt.Sprintf("(*%s, error)", m.ResponseType)
			} else {
				fnSig += "error"
			}
			lines = append(lines, fnSig+" {")
			if m.ResponseType != "" {
				lines = append(lines, fmt.Sprintf("\tvar resp %s", m.ResponseType))
			}
			method := "Post"
			if m.HTTPMethod == "GET" {
				method = "Get"
			}
			httpCall := fmt.Sprintf("\t_, err := s.t.%s(ctx, %q", method, m.Path)
			if m.RequestType != "" {
				if m.ResponseType != "" {
					httpCall += ", req, &resp)"
				} else {
					httpCall += ", req, nil)"
				}
			} else {
				if m.ResponseType != "" {
					httpCall += ", nil, &resp)"
				} else {
					httpCall += ", nil, nil)"
				}
			}
			lines = append(lines, httpCall)
			lines = append(lines, "\tif err != nil {")
			if m.ResponseType != "" {
				lines = append(lines, "\t\treturn nil, err")
			} else {
				lines = append(lines, "\t\treturn err")
			}
			lines = append(lines, "\t}")
			if m.ResponseType != "" {
				lines = append(lines, "\treturn &resp, nil")
			} else {
				lines = append(lines, "\treturn nil")
			}
			lines = append(lines, "}", "")
		}
		os.WriteFile(implFile, []byte(strings.Join(lines, "\n")), 0644)
		fmt.Printf("%s: %d methods\n", implFile, len(mlist))
	}
}
