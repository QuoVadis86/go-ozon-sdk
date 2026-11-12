//go:build ignore

package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGeneratedOrderKeepsExistingDeclarations(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "types.go"), []byte("package example\n"+
		"type Zebra struct {\n"+
		"Second string `json:\"second\"`\n"+
		"First string `json:\"first\"`\n"+
		"}\n"+
		"type Alpha struct{}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "service.go"), []byte(`package example
type Service struct{}
func (s *Service) Old() {}
`), 0644); err != nil {
		t.Fatal(err)
	}

	layout := readGeneratedLayout(dir)
	if got, want := orderedKeys(map[string]bool{"first": true, "second": true, "added_b": true, "added_a": true}, layout.fields["Zebra"]), []string{"second", "first", "added_a", "added_b"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("field order = %v, want %v", got, want)
	}
	types := []string{"NewZ", "Alpha", "NewA", "Zebra"}
	orderByExisting(types, func(name string) string { return name }, layout.types)
	if want := []string{"Zebra", "Alpha", "NewA", "NewZ"}; !reflect.DeepEqual(types, want) {
		t.Fatalf("type order = %v, want %v", types, want)
	}
	methods := []string{"NewZ", "Old", "NewA"}
	orderByExisting(methods, func(name string) string { return name }, layout.methods)
	if want := []string{"Old", "NewA", "NewZ"}; !reflect.DeepEqual(methods, want) {
		t.Fatalf("method order = %v, want %v", methods, want)
	}
}

func TestResolveFieldType_ArrayWithSiblingRef(t *testing.T) {
	cache := map[string]string{"DetailsPayment": "DetailsPayment"}
	for _, tc := range []struct {
		name  string
		field map[string]interface{}
		want  string
	}{
		{"array with null items and sibling ref", map[string]interface{}{"type": "array", "items": nil, "$ref": "#/components/schemas/DetailsPayment"}, "[]DetailsPayment"},
		{"scalar ref", map[string]interface{}{"$ref": "#/components/schemas/DetailsPayment"}, "DetailsPayment"},
		{"array with item ref", map[string]interface{}{"type": "array", "items": map[string]interface{}{"$ref": "#/components/schemas/DetailsPayment"}}, "[]DetailsPayment"},
		{"array without item schema", map[string]interface{}{"type": "array", "items": nil}, "[]any"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := resolveFieldType(tc.field, cache, nil); got != tc.want {
				t.Fatalf("resolveFieldType() = %q, want %q", got, tc.want)
			}
		})
	}

	if got := resolveFieldTypeFromSchema(&Schema{Type: "array", Ref: "#/components/schemas/DetailsPayment"}, cache, nil); got != "[]DetailsPayment" {
		t.Fatalf("resolveFieldTypeFromSchema() = %q, want []DetailsPayment", got)
	}
}
