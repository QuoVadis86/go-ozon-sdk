package ozon

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestDeprecatedMethodsHaveGoDocNotices(t *testing.T) {
	methods := map[string][]string{
		"fbs/service.go":                  {"GetFbsPostingListV3", "GetFbsPostingUnfulfilledList", "PostingFBSPackageLabel"},
		"finance/service.go":              {"FinanceTransactionListV3", "FinanceTransactionTotalV3"},
		"product/service.go":              {"ProductImportPictures"},
		"product/attributes_v3_compat.go": {"GetProductAttributesV3"},
		"promo/service.go": {
			"ActionsAutoAddProductsCandidates", "ActionsAutoAddProductsDelete",
			"ActionsAutoAddProductsList", "ActionsAutoAddProductsUpdate",
			"PromosCandidates", "PromosProducts", "PromosProductsActivate",
			"PromosProductsDeactivate", "TaskList",
		},
		"review/service.go": {"ReviewList"},
		"returns/legacy_rfbs.go": {
			"ReturnsRfbsVerifyV2", "ReturnsRfbsReturnMoneyV2", "ReturnsRfbsCompensateV2",
			"ReturnsRfbsRejectV2", "ReturnsRfbsReceiveReturnV2",
		},
	}
	for path, names := range methods {
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		found := map[string]bool{}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				continue
			}
			for _, name := range names {
				if fn.Name.Name != name {
					continue
				}
				for _, paragraph := range strings.Split(fn.Doc.Text(), "\n\n") {
					if strings.HasPrefix(paragraph, "Deprecated: ") {
						found[name] = true
					}
				}
			}
		}
		for _, name := range names {
			if !found[name] {
				t.Errorf("%s: %s has no standalone Deprecated: paragraph", path, name)
			}
		}
	}
}
