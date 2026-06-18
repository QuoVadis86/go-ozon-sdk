package category

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"` // 类别清单。
}

type V1GetAttributesRequest struct {
	Language interface{} `json:"language"`
	TypeID int64 `json:"type_id"` // 商品类型ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类别ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"` // 请求结果。
}

type V1GetAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"` // 特性ID。可以使用方法 [/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes)获取。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类别ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	Language interface{} `json:"language"`
	LastValueID int64 `json:"last_value_id"` // 启动响应的指南 ID。 如果`last_value_id`为 10，则响应将包含从第十一个开始的指南。
	Limit int64 `json:"limit"` // 响应中值的数量： - 最多 —— 2000， - 最少 —— 1。
	TypeID int64 `json:"type_id"` // 商品类型ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1GetAttributeValuesResponse struct {
	HasNext bool `json:"has_next"` // 该特征表示响应中只返回了部分特性值： - `true` —— 请用新参数 `last_value_id` 再次请求以获取其它值； - `false` —— 响应包含了所有特性值。
	Result []interface{} `json:"result"` // 特性值。
}

type V1SearchAttributeValuesRequest struct {
	Limit int64 `json:"limit"` // 返回结果中的值数量：: - 最大值 — 100， - 最小值 — 1。
	TypeID int64 `json:"type_id"` // 商品类型的标识符。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
	Value string `json:"value"` // 系统将根据此值搜索参考值。最少需要2个字符。
	AttributeID int64 `json:"attribute_id"` // 属性的标识符。可以使用方法 [/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes)获取。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类目的标识符。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"` // 属性值。
}

type V1GetTreeRequest struct {
	Language interface{} `json:"language"`
}
