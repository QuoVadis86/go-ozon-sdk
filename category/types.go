package category

type V1GetTreeResponseItem struct {
	TypeID                int64                   `json:"type_id"`                 // 商品类型ID。
	TypeName              string                  `json:"type_name"`               // 商品类型名称。
	DescriptionCategoryID int64                   `json:"description_category_id"` // 类别ID。
	CategoryName          string                  `json:"category_name"`           // 类别名称。
	Children              []V1GetTreeResponseItem `json:"children"`                // 子类别树形图。
	Disabled              bool                    `json:"disabled"`                // 如果无法在类别中创建商品，则为`true`。 如果可能，为`false`。
}

type V1GetTreeResponse struct {
	Result []V1GetTreeResponseItem `json:"result"` // 类别清单。
}

// 商品属性的值。
type SearchAttributeValuesResponseValue struct {
	Info    string `json:"info"`    // 额外信息。
	Picture string `json:"picture"` // 图片链接。
	Value   string `json:"value"`
	ID      int64  `json:"id"` // 属性值的标识符。
}

// 回复语言： - `EN` — 英语， - `RU` — 俄语， - `TR` — 土耳其语， - `ZH_HANS` — 中文。 默认情况下，使用俄语。
type LanguageLanguage string

type V1GetAttributesRequest struct {
	DescriptionCategoryID int64            `json:"description_category_id"` // 类别ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	Language              LanguageLanguage `json:"language"`
	TypeID                int64            `json:"type_id"` // 商品类型ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1GetAttributesResponseAttribute struct {
	ID                  int64  `json:"id"`                    // 特性ID。
	IsAspect            bool   `json:"is_aspect"`             // 方面属性特征。方面属性：用于区分同类商品不同特征的属性。 例如，同款衣服和鞋子具有不同的颜色和尺寸，即：颜色、尺寸为方面属性。 字段值： - `true` — 方面属性，在货物交付到仓库或出仓销售后不能更改。 - `false` — 非方面...
	Type                string `json:"type_"`                 // 特征类型。
	AttributeComplexID  int64  `json:"attribute_complex_id"`  // 复合属性的标识符。
	CategoryDependent   bool   `json:"category_dependent"`    // 字典属性值取决于类别的标志： - `true` — 该属性对每个类别都有不一样的值。 - `false` — 该属性对所有类别都有相同的值。
	GroupID             int64  `json:"group_id"`              // 组别特征ID。
	GroupName           string `json:"group_name"`            // 特征组别名称。
	IsCollection        bool   `json:"is_collection"`         // 标志、特征 — 值集： - `true` — 特征 — 值集, - `false` — 特性由单个值组成。
	IsRequired          bool   `json:"is_required"`           // 强制性特征标志: - `true` — 强制性特征, - `false` — 可不指出特征。
	Name                string `json:"name"`                  // 名称。
	MaxValueCount       int64  `json:"max_value_count"`       // 属性的最大值数量。
	ComplexIsCollection bool   `json:"complex_is_collection"` // 标志某个复合特征是否为值集合： - `true` 表示该复合特征是一个值集合； - `false` 表示该复合特征只有一个值。
	Description         string `json:"description"`           // 特征描述。
	DictionaryID        int64  `json:"dictionary_id"`         // 目录ID。
}

type V1GetAttributesResponse struct {
	Result []V1GetAttributesResponseAttribute `json:"result"` // 请求结果。
}

type V1GetAttributeValuesRequest struct {
	AttributeID           int64            `json:"attribute_id"`            // 特性ID。可以使用方法 [/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes)获取。
	DescriptionCategoryID int64            `json:"description_category_id"` // 类别ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	Language              LanguageLanguage `json:"language"`
	LastValueID           int64            `json:"last_value_id"` // 启动响应的指南 ID。 如果`last_value_id`为 10，则响应将包含从第十一个开始的指南。
	Limit                 int64            `json:"limit"`         // 响应中值的数量： - 最多 —— 2000， - 最少 —— 1。
	TypeID                int64            `json:"type_id"`       // 商品类型ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1GetAttributeValuesResponseDictionaryValue struct {
	ID      int64  `json:"id"`      // 特性值ID。
	Info    string `json:"info"`    // 附加描述。
	Picture string `json:"picture"` // 图片链接。
	Value   string `json:"value"`   // 商品特性值。
}

type V1GetAttributeValuesResponse struct {
	Result  []V1GetAttributeValuesResponseDictionaryValue `json:"result"`   // 特性值。
	HasNext bool                                          `json:"has_next"` // 该特征表示响应中只返回了部分特性值： - `true` —— 请用新参数 `last_value_id` 再次请求以获取其它值； - `false` —— 响应包含了所有特性值。
}

type V1SearchAttributeValuesRequest struct {
	Limit                 int64  `json:"limit"`                   // 返回结果中的值数量：: - 最大值 — 100， - 最小值 — 1。
	TypeID                int64  `json:"type_id"`                 // 商品类型的标识符。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	Value                 string `json:"value"`                   // 系统将根据此值搜索参考值。最少需要2个字符。
	AttributeID           int64  `json:"attribute_id"`            // 属性的标识符。可以使用方法 [/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes)获取。
	DescriptionCategoryID int64  `json:"description_category_id"` // 类目的标识符。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1SearchAttributeValuesResponse struct {
	Result []SearchAttributeValuesResponseValue `json:"result"` // 属性值。
}

type V1GetTreeRequest struct {
	Language LanguageLanguage `json:"language"`
}
