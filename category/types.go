package category

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetAttributeValuesRequest struct {
	Language interface{} `json:"language"`
	LastValueID int64 `json:"last_value_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
}

type V1GetAttributeValuesResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1SearchAttributeValuesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	Value string `json:"value"`
	AttributeID int64 `json:"attribute_id"`
}

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetTreeRequest struct {
	Language interface{} `json:"language"`
}

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetAttributesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	TypeID int64 `json:"type_id"`
}
