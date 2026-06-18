//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
		"ReviewAPI":          "review", "Questions&Answers": "review",
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

type Schema struct {
	Type        string                 `json:"type"`
	Properties  map[string]interface{} `json:"properties"`
	Items       *Schema                `json:"items"`
	Enum        []interface{}          `json:"enum"`
	Description string                 `json:"description"`
	Title       string                 `json:"title"`
	Ref         string                 `json:"$ref"`
	Format      string                 `json:"format"`
	Nullable    bool                   `json:"nullable"`
}

// resolveFieldType determines the Go type for a property, handling enums, arrays, etc.
func resolveFieldType(p map[string]interface{}, cache map[string]string, schemas map[string]Schema) string {
	ref, _ := p["$ref"].(string)
	if ref != "" {
		rname := strings.TrimPrefix(ref, "#/components/schemas/")
		if t, ok := cache[rname]; ok {
			return t
		}
		return "any"
	}

	typ, _ := p["type"].(string)
	fmt := ""
	if f, ok := p["format"].(string); ok {
		fmt = f
	}

	// Check for array via items field (even without explicit type: array)
	if _, hasItems := p["items"]; hasItems {
		itemsRaw, _ := json.Marshal(p["items"])
		var items Schema
		json.Unmarshal(itemsRaw, &items)
		itemType := resolveFieldTypeFromSchema(&items, cache, schemas)
		return "[]" + itemType
	}

	switch typ {
	case "array":
		return "[]any"
	case "string":
		return "string"
	case "integer":
		if fmt == "int32" {
			return "int32"
		}
		return "int64"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "object":
		return "map[string]any"
	default:
		return "any"
	}
}

func resolveFieldTypeFromSchema(s *Schema, cache map[string]string, schemas map[string]Schema) string {
	if s.Ref != "" {
		rname := strings.TrimPrefix(s.Ref, "#/components/schemas/")
		if t, ok := cache[rname]; ok {
			return t
		}
		return "any"
	}
	if s.Type == "array" || s.Items != nil {
		sub := resolveFieldTypeFromSchema(s.Items, cache, schemas)
		return "[]" + sub
	}
	switch s.Type {
	case "string":
		return "string"
	case "integer":
		if s.Format == "int32" {
			return "int32"
		}
		return "int64"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "object":
		return "map[string]any"
	default:
		return "any"
	}
}

func containsAny(s string, list []string) bool {
	for _, item := range list {
		if s == item {
			return true
		}
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

var enumRe = regexp.MustCompile("`([^`]+)`")

var descRe = regexp.MustCompile("`([^`]+)`(?:\\s*[-–—]+\\s*([^\n\r;]*))?")

// extractEnumValues parses field description for inline enum values like:
// - `value` — description
// Returns deduplicated non-boolean string values.
func extractEnumValues(desc string) []string {
	vals, _ := extractEnumValuesWithDesc(desc)
	return vals
}

// extractEnumValuesWithDesc returns values and their descriptions.
func extractEnumValuesWithDesc(desc string) ([]string, []string) {
	if desc == "" {
		return nil, nil
	}
	matches := descRe.FindAllStringSubmatch(desc, -1)
	if len(matches) < 2 {
		return nil, nil
	}
	boolSet := map[string]bool{"true": true, "false": true, "0": true, "1": true, "yes": true, "no": true}
	seen := map[string]bool{}
	var vals []string
	var valDescs []string
	for _, m := range matches {
		v := strings.TrimSpace(m[1])
		if v == "" || boolSet[strings.ToLower(v)] {
			continue
		}
		if seen[v] {
			continue
		}
		seen[v] = true
		vals = append(vals, v)
		d := strings.TrimSpace(m[2])
		valDescs = append(valDescs, d)
	}
	if len(vals) < 2 {
		return nil, nil
	}
	return vals, valDescs
}

// shortName shortens a struct+field name to a minimal but meaningful type name.
// It uses just the field name, with struct context only when needed for uniqueness.
func shortName(structName, fieldName string) string {
	fn := toGoName(fieldName)
	// If field name is already distinctive enough, use it alone
	if len(fn) >= 3 {
		return fn
	}
	// Otherwise combine with the last meaningful part of struct name
	parts := splitCamel(structName)
	if len(parts) > 1 {
		return parts[len(parts)-1] + fn
	}
	return fn
}

func splitCamel(s string) []string {
	var parts []string
	var cur []rune
	for _, r := range s {
		if r >= 'A' && r <= 'Z' && len(cur) > 0 {
			parts = append(parts, string(cur))
			cur = []rune{r}
		} else {
			cur = append(cur, r)
		}
	}
	if len(cur) > 0 {
		parts = append(parts, string(cur))
	}
	return parts
}

// shortEnumTypeName creates a short Go type name for an enum.
func shortEnumTypeName(structName, fieldName string, usedNames map[string]string) string {
	base := shortName(structName, fieldName)
	if _, exists := usedNames[base]; !exists {
		return base
	}
	// Collision: prepend distinguishing part from struct name
	sParts := splitCamel(toGoName(structName))
	for i := len(sParts) - 1; i >= 0; i-- {
		candidate := sParts[i] + base
		if _, exists := usedNames[candidate]; !exists {
			return candidate
		}
	}
	return toGoName(structName + "_" + fieldName)
}

// cleanEnumValue sanitizes a raw enum value string into a Go identifier fragment.
func cleanEnumValue(value string) string {
	s := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		return '_'
	}, value)
	s = strings.Trim(s, "_")
	if s == "" || (s[0] >= '0' && s[0] <= '9') {
		s = "V" + s
	}
	return s
}

// valueToCamel converts a raw value like "FBS_DELIVERY" to "FBSDelivery" or "pending" to "Pending".
func valueToCamel(value string) string {
	clean := cleanEnumValue(value)
	parts := strings.Split(clean, "_")
	var cased []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		allUpper := strings.ToUpper(p) == p
		if len(p) <= 3 && allUpper {
			// Short all-caps like FBS, ID -> keep uppercase
			cased = append(cased, p)
		} else if allUpper {
			// Long all-caps like DELIVERY -> capitalize first letter
			cased = append(cased, strings.ToUpper(p[:1])+strings.ToLower(p[1:]))
		} else {
			cased = append(cased, strings.ToUpper(p[:1])+p[1:])
		}
	}
	return strings.Join(cased, "")
}

var (
	enumTypeGenerated map[string]bool
	allGeneratedTypes map[string]bool
	enumNameUsed      map[string]string
)

func appendType(out *[]string, name string, s Schema, cache map[string]string, schemas map[string]Schema) {
	if s.Description != "" {
		*out = append(*out, fmt.Sprintf("// %s", sanitizeComment(s.Description)))
	}
	if len(s.Properties) == 0 {
		if allGeneratedTypes[name] {
			*out = append(*out, "")
			return
		}
		allGeneratedTypes[name] = true
		switch s.Type {
		case "array":
			itemType := "any"
			if s.Items != nil {
				itemType = resolveFieldTypeFromSchema(s.Items, cache, schemas)
			}
			*out = append(*out, fmt.Sprintf("type %s []%s", name, itemType))
		case "string":
			*out = append(*out, fmt.Sprintf("type %s string", name))
		case "integer":
			*out = append(*out, fmt.Sprintf("type %s int64", name))
		case "number":
			*out = append(*out, fmt.Sprintf("type %s float64", name))
		case "boolean":
			*out = append(*out, fmt.Sprintf("type %s bool", name))
		default:
			*out = append(*out, fmt.Sprintf("type %s any", name))
		}
		*out = append(*out, "")
		return
	}

	// First pass: collect inline enum definitions and generate type+const blocks
	type enumDef struct {
		typeName   string
		values     []string
		valueDescs []string
	}
	var enums []enumDef
	for pn, pv := range s.Properties {
		pvMap, _ := pv.(map[string]interface{})
		if pvMap == nil {
			continue
		}
		desc, _ := pvMap["description"].(string)
		vals, valDescs := extractEnumValuesWithDesc(desc)
		if vals == nil {
			continue
		}
		typeName := shortEnumTypeName(name, pn, enumNameUsed)
		enums = append(enums, enumDef{
			typeName:   typeName,
			values:     vals,
			valueDescs: valDescs,
		})
	}

	// Generate enum types and constants before the struct
	for _, e := range enums {
		if allGeneratedTypes[e.typeName] {
			continue
		}
		allGeneratedTypes[e.typeName] = true
		*out = append(*out, fmt.Sprintf("// %s values", e.typeName))
		*out = append(*out, fmt.Sprintf("type %s string", e.typeName))
		*out = append(*out, "const (")
		usedNames := map[string]bool{}
		for i, v := range e.values {
			constName := e.typeName + valueToCamel(v)
			if usedNames[constName] {
				suffix := 1
				for usedNames[fmt.Sprintf("%s_%d", constName, suffix)] {
					suffix++
				}
				constName = fmt.Sprintf("%s_%d", constName, suffix)
			}
			usedNames[constName] = true
			line := fmt.Sprintf("\t%s %s = %q", constName, e.typeName, v)
			if i < len(e.valueDescs) && e.valueDescs[i] != "" {
				line += fmt.Sprintf(" // %s", sanitizeComment(e.valueDescs[i]))
			}
			*out = append(*out, line)
		}
		*out = append(*out, ")")
		*out = append(*out, "")
	}

	if allGeneratedTypes[name] {
		*out = append(*out, "")
		return
	}
	allGeneratedTypes[name] = true
	*out = append(*out, fmt.Sprintf("type %s struct {", name))
	for pn, pv := range s.Properties {
		pvMap, _ := pv.(map[string]interface{})
		if pvMap == nil {
			continue
		}
		desc, _ := pvMap["description"].(string)
		fn := toGoName(pn)
		ft := resolveFieldType(pvMap, cache, schemas)
		if ft == "string" && extractEnumValues(desc) != nil {
			ft = shortEnumTypeName(name, pn, enumNameUsed)
		}
		jtag := pn
		if pn == "type" {
			jtag = "type_"
		}
		line := fmt.Sprintf("\t%s %s `json:\"%s\"`", fn, ft, jtag)
		if desc != "" {
			comment := sanitizeComment(desc)
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
	if err != nil {
		panic(err)
	}
	var raw map[string]json.RawMessage
	json.Unmarshal(data, &raw)
	var paths map[string]map[string]json.RawMessage
	json.Unmarshal(raw["paths"], &paths)
	var comps struct {
		Schemas map[string]json.RawMessage `json:"schemas"`
	}
	json.Unmarshal(raw["components"], &comps)

	schemas := map[string]Schema{}
	for n, raw := range comps.Schemas {
		var s Schema
		json.Unmarshal(raw, &s)
		schemas[n] = s
	}
	cache := map[string]string{}
	svcPrefixes := []string{"product", "finance", "warehouse", "seller", "fbs", "fbo", "promo", "beta", "premium", "pass", "returns", "review", "chat", "barcode", "rating", "delivery", "report", "category", "certificate"}
	for n := range comps.Schemas {
		stripped := n
		lower := strings.ToLower(n)
		for _, p := range svcPrefixes {
			ln := len(p)
			if len(stripped) <= ln {
				continue
			}
			if lower[:ln] == p && (lower[ln:ln+1] == "." || (lower[ln] >= 'a' && lower[ln] <= 'z')) {
				// Strip if followed by dot or another lowercase letter (camelCase continuation)
				// For camelCase: productV3Xxx -> V3Xxx (strip product, keep V3)
				stripped = n[ln:]
				if lower[ln:ln+1] == "." {
					stripped = n[ln+1:]
				}
				break
			}
		}
		cache[n] = toCamel(stripped)
	}

	depOf := map[string]map[string]bool{}
	for sname, s := range schemas {
		depOf[sname] = map[string]bool{}
		if len(s.Properties) == 0 {
			continue
		}
		for _, pv := range s.Properties {
			pr, _ := json.Marshal(pv)
			var p struct {
				Ref   string `json:"$ref"`
				Items *struct {
					Ref string `json:"$ref"`
				} `json:"items"`
			}
			json.Unmarshal(pr, &p)
			if p.Ref != "" {
				rn := strings.TrimPrefix(p.Ref, "#/components/schemas/")
				if rn != sname {
					depOf[sname][rn] = true
				}
			}
			if p.Items != nil && p.Items.Ref != "" {
				rn := strings.TrimPrefix(p.Items.Ref, "#/components/schemas/")
				if rn != sname {
					depOf[sname][rn] = true
				}
			}
		}
	}

	type Method struct {
		Dir, Name, HTTPM, Path, ReqT, RespT, Summary string
		Deprecated                                   bool
		ReplaceWith                                  string
		Notes                                        []string
	}
	dirMethods := map[string][]Method{}
	dirDirectTypes := map[string]map[string]bool{}
	methodNames := map[string]bool{}

	for path, items := range paths {
		for httpM, raw := range items {
			var item struct {
				Summary     string                     `json:"summary"`
				Description string                     `json:"description"`
				OperationID string                     `json:"operationId"`
				Tags        []string                   `json:"tags"`
				Deprecated  bool                       `json:"deprecated"`
				RequestBody *json.RawMessage           `json:"requestBody"`
				Responses   map[string]json.RawMessage `json:"responses"`
			}
			json.Unmarshal(raw, &item)
			dir := ""
			for _, t := range item.Tags {
				if d, ok := tagDir[t]; ok {
					dir = d
					break
				}
			}
			if dir == "" {
				continue
			}
			if dirDirectTypes[dir] == nil {
				dirDirectTypes[dir] = map[string]bool{}
			}

			reqT := ""
			if item.RequestBody != nil {
				var rb struct {
					Content map[string]struct {
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema"`
					} `json:"content"`
				}
				json.Unmarshal(*item.RequestBody, &rb)
				for _, c := range rb.Content {
					if c.Schema.Ref != "" {
						sname := strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")
						reqT = cache[sname]
						dirDirectTypes[dir][sname] = true
						break
					}
				}
			}
			respT := ""
			if r, ok := item.Responses["200"]; ok {
				var rb struct {
					Content map[string]struct {
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema"`
					} `json:"content"`
				}
				json.Unmarshal(r, &rb)
				for _, c := range rb.Content {
					if c.Schema.Ref != "" {
						sname := strings.TrimPrefix(c.Schema.Ref, "#/components/schemas/")
						respT = cache[sname]
						dirDirectTypes[dir][sname] = true
						break
					}
				}
			}

			name := item.OperationID
			parts := strings.Split(name, "_")
			if len(parts) > 1 {
				name = toGoName(strings.Join(parts[1:], "_"))
			} else {
				name = toGoName(name)
			}

			k := dir + ":" + name
			if _, exists := methodNames[k]; exists {
				for _, p := range strings.Split(strings.Trim(path, "/"), "/") {
					if !strings.ContainsAny(p, "{}") {
						name += toGoName(p)
						break
					}
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
				if idx < 0 {
					continue
				}
				start := idx
				end := idx + 130
				if end > len(desc) {
					end = len(desc)
				}
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
			if visited[sname] {
				return
			}
			visited[sname] = true
			for dep := range depOf[sname] {
				visit(dep)
			}
		}
		for s := range direct {
			visit(s)
		}
		dirAllTypes[dir] = visited
	}

	for dir, methods := range dirMethods {
		dp := filepath.Join(dir)
		os.MkdirAll(dp, 0755)

		dlines := []string{fmt.Sprintf("package %s", dir), ""}
		typeCount := 0
		dirGenned := map[string]bool{}
		dirInProg := map[string]bool{}
		dirGoNames := map[string]string{}
		enumTypeGenerated = map[string]bool{}
		allGeneratedTypes = map[string]bool{}
		enumNameUsed = map[string]string{}

		var writeType func(string)
		writeType = func(sname string) {
			if dirGenned[sname] || dirInProg[sname] {
				return
			}
			dirInProg[sname] = true
			s, ok := schemas[sname]
			if !ok {
				dirInProg[sname] = false
				return
			}
			gn := cache[sname]
			if prev, exists := dirGoNames[gn]; exists && prev != sname {
				dirInProg[sname] = false
				dirGenned[sname] = true
				return
			}
			dirGoNames[gn] = sname
			for dep := range depOf[sname] {
				if dep != sname {
					writeType(dep)
				}
			}
			dirInProg[sname] = false
			dirGenned[sname] = true
			if dirAllTypes[dir][sname] {
				appendType(&dlines, gn, s, cache, schemas)
				typeCount++
			}
		}

		for t := range dirAllTypes[dir] {
			writeType(t)
		}

		os.WriteFile(dp+"/types.go", []byte(strings.Join(dlines, "\n")), 0644)

		var ml []string
		ml = append(ml, fmt.Sprintf("package %s", dir), "")
		imports := `"context"; "github.com/QuoVadis86/go-ozon-sdk/transport"`
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
			if m.ReqT != "" {
				fn += fmt.Sprintf(", req *%s", m.ReqT)
			}
			fn += ") "
			if m.RespT != "" {
				fn += fmt.Sprintf("(*%s, error)", m.RespT)
			} else {
				fn += "error"
			}
			ml = append(ml, fn+" {")
			if m.RespT != "" {
				ml = append(ml, fmt.Sprintf("\tvar resp %s", m.RespT))
			}
			meth := "Post"
			if m.HTTPM == "GET" {
				meth = "Get"
			}
			call := fmt.Sprintf("\terr := s.Client.%s(ctx, %q", meth, m.Path)
			if m.ReqT != "" {
				if m.RespT != "" {
					call += ", req, &resp)"
				} else {
					call += ", req, nil)"
				}
			} else {
				if m.RespT != "" {
					call += ", nil, &resp)"
				} else {
					call += ", nil, nil)"
				}
			}
			ml = append(ml, call)
			ml = append(ml, "\tif err != nil {")
			if m.RespT != "" {
				ml = append(ml, "\t\treturn nil, err")
			} else {
				ml = append(ml, "\t\treturn err")
			}
			ml = append(ml, "\t}")
			if m.RespT != "" {
				ml = append(ml, "\treturn &resp, nil")
			} else {
				ml = append(ml, "\treturn nil")
			}
			ml = append(ml, "}", "")
		}
		os.WriteFile(dp+"/service.go", []byte(strings.Join(ml, "\n")), 0644)

		// Generate tests
		tlines := []string{
			fmt.Sprintf("package %s", dir), "",
			`import ("context"; "testing"; "github.com/QuoVadis86/go-ozon-sdk/transport")`, "",
			"var ctx = context.Background()", "",
		}
		// Find first method with valid request+response types for a basic test
		var testMethod Method
		skipTypes := map[string]bool{"V1Empty": true, "Polygonv1Empty": true, "CreateDiscountedRequest": true}
		for _, m := range methods {
			if m.RespT != "" && !skipTypes[m.RespT] && !skipTypes[m.ReqT] {
				testMethod = m
				break
			}
		}
		if testMethod.Name != "" {
			testName := "Test" + testMethod.Name
			tlines = append(tlines, fmt.Sprintf("func %s(t *testing.T) {", testName))
			tlines = append(tlines, "\thandler := transport.MockHandler(200, "+testMethod.RespT+"{})")
			tlines = append(tlines, "\tcl, srv := transport.NewTestClient(handler)")
			tlines = append(tlines, "\tdefer srv.Close()")
			tlines = append(tlines, "\tsvc := &Service{Client: cl}")
			call := fmt.Sprintf("\tresp, err := svc.%s(ctx", testMethod.Name)
			if testMethod.ReqT != "" {
				call += ", &" + testMethod.ReqT + "{}"
			}
			call += ")"
			tlines = append(tlines, call)
			tlines = append(tlines, "\tif err != nil {")
			tlines = append(tlines, fmt.Sprintf("\t\tt.Fatalf(\"%s() error: %%v\", err)", testMethod.Name))
			tlines = append(tlines, "\t}")
			tlines = append(tlines, "\tif resp == nil {")
			tlines = append(tlines, fmt.Sprintf("\t\tt.Fatal(\"%s() returned nil\")", testMethod.Name))
			tlines = append(tlines, "\t}")
			tlines = append(tlines, "}", "")
		}
		// Error handling test
		if testMethod.Name != "" {
			tlines = append(tlines, "func TestAPIError(t *testing.T) {")
			tlines = append(tlines, "\thandler := transport.MockHandler(400, map[string]interface{}{")
			tlines = append(tlines, `		"code": 400,`)
			tlines = append(tlines, `		"message": "test error",`)
			tlines = append(tlines, "\t})")
			tlines = append(tlines, "\tcl, srv := transport.NewTestClient(handler)")
			tlines = append(tlines, "\tdefer srv.Close()")
			tlines = append(tlines, "\tsvc := &Service{Client: cl}")
			call := fmt.Sprintf("\t_, err := svc.%s(ctx", testMethod.Name)
			if testMethod.ReqT != "" {
				call += ", &" + testMethod.ReqT + "{}"
			}
			call += ")"
			tlines = append(tlines, call)
			tlines = append(tlines, "\tif err == nil {")
			tlines = append(tlines, "\t\tt.Fatal(\"expected error, got nil\")")
			tlines = append(tlines, "\t}")
			tlines = append(tlines, "}", "")
		}
		os.WriteFile(dp+"/service_test.go", []byte(strings.Join(tlines, "\n")), 0644)

		fmt.Printf("%s: %d methods, %d types\n", dir, len(methods), typeCount)
	}

	// Format all generated files
	exec.Command("go", "fmt", "./...").Run()
}
