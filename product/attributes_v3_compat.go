package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/enum"
)

type GetProductAttributesResponsePdf struct {
	Name     string `json:"name"`      // 文件名称。
	FileName string `json:"file_name"` // 到PDF文件的路径。
	Index    int64  `json:"index"`     // 存储库中的设定了顺序的文件指数。
}

type V3Filter struct {
	OfferID    []string                                      `json:"offer_id"`   // 基于参数 `offer_id`的过滤。 可以提交数值列表。
	ProductID  []string                                      `json:"product_id"` // 基于参数 `product_id`的过滤。 可以提交数值列表。
	Visibility V2GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type V3GetProductAttributesV3Request struct {
	SortBy  string   `json:"sort_by"`  // 对商品进行分类的参数。
	SortDir string   `json:"sort_dir"` // 分类方向。
	Filter  V3Filter `json:"filter"`
	LastID  string   `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit   int64    `json:"limit"`   // 每页的值的数量。最小 —— 1，最大 —— 1000。
}

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64                                           `json:"attribute_id"` // 特征的识别码。
	ComplexID   int64                                           `json:"complex_id"`   // 支持嵌入属性的特征的识别码。例如，"处理器" 有嵌入属性的特征 "制造商" 和 "二级缓存"。每个嵌入属性都可以有几个值选项。
	Values      []GetProductAttributesV3ResponseDictionaryValue `json:"values"`       // 特征值的数组。
}

type GetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []GetProductAttributesV3ResponseAttribute `json:"attributes"` // 商品特征的数组。
}

type GetProductAttributesResponseImage360 struct {
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
}

type V3GetProductAttributesV3ResponseResult struct {
	CategoryID            int64                                            `json:"category_id"`             // 类别ID。 请将其与以下方法结合使用：[/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree), [/v1/description-category...
	Images360             []GetProductAttributesResponseImage360           `json:"images360"`               // 图像数组360。
	OfferID               string                                           `json:"offer_id"`                // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	PDFList               []GetProductAttributesResponsePdf                `json:"pdf_list"`                // PDF文件的阵列。
	Attributes            []GetProductAttributesV3ResponseAttribute        `json:"attributes"`              // 商品特性的数组。
	Height                int32                                            `json:"height"`                  // 包装高度。
	ID                    int64                                            `json:"id"`                      // 商品特性的识别码。
	Width                 int32                                            `json:"width"`                   // 包装宽度。
	Barcode               string                                           `json:"barcode"`                 // 条形码。
	DescriptionCategoryID int64                                            `json:"description_category_id"` // 类别ID。 请将其与以下方法结合使用：[/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes) 和 [/v1/descript...
	Depth                 int32                                            `json:"depth"`                   // 深度。
	Images                []GetProductAttributesResponseImage              `json:"images"`                  // 产品图片链接的数组。
	Weight                int32                                            `json:"weight"`                  // 商品在包装中的重量。
	TypeID                int64                                            `json:"type_id"`                 // 商品类型的标识符。
	ColorImage            string                                           `json:"color_image"`             // 市场营销色彩。
	ComplexAttributes     []GetProductAttributesV3ResponseComplexAttribute `json:"complex_attributes"`      // 已录入的特性的数组。
	DimensionUnit         enum.DimensionUnit                               `json:"dimension_unit"`          // 尺寸的测量单位。 - `mm` —— 毫米， - `cm` —— 厘米， - `in` —— 英寸。
	ImageGroupID          string                                           `json:"image_group_id"`          // 用于后续批量下载图像的识别码。
	Name                  string                                           `json:"name"`                    // 商品名称。最多500个字符。
	WeightUnit            string                                           `json:"weight_unit"`             // 重量测量单位。
}

type V3GetProductAttributesV3Response struct {
	Result []V3GetProductAttributesV3ResponseResult `json:"result"`  // 查询结果。
	LastID string                                   `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Total  string                                   `json:"total"`   // 列表中的商品数量。
}

// GetProductAttributesV3 is retained for callers using the retired v3 API.
// Deprecated: use GetProductAttributesV4.
func (s *Service) GetProductAttributesV3(ctx context.Context, req *V3GetProductAttributesV3Request) (*V3GetProductAttributesV3Response, error) {
	var resp V3GetProductAttributesV3Response
	if err := s.Client.Post(ctx, "/v4/product/info/attributes", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
