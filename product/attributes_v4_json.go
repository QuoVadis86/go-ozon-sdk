package product

import "encoding/json"

func (image *GetProductAttributesResponseImage) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var url string
		if err := json.Unmarshal(data, &url); err != nil {
			return err
		}
		*image = GetProductAttributesResponseImage{URL: url}
		return nil
	}

	type imageAlias GetProductAttributesResponseImage
	var decoded imageAlias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*image = GetProductAttributesResponseImage(decoded)
	return nil
}

func (result *V4GetProductAttributesV4ResponseResult) UnmarshalJSON(data []byte) error {
	type resultAlias V4GetProductAttributesV4ResponseResult
	decoded := struct {
		SKU json.RawMessage `json:"sku"`
		*resultAlias
	}{resultAlias: (*resultAlias)(result)}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	return unmarshalStringOrNumber(decoded.SKU, &result.SKU)
}

func (response *V4GetProductAttributesV4Response) UnmarshalJSON(data []byte) error {
	type responseAlias V4GetProductAttributesV4Response
	decoded := struct {
		Total json.RawMessage `json:"total"`
		*responseAlias
	}{responseAlias: (*responseAlias)(response)}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	return unmarshalStringOrNumber(decoded.Total, &response.Total)
}

func unmarshalStringOrNumber(data json.RawMessage, value *string) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		return json.Unmarshal(data, value)
	}
	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}
	*value = number.String()
	return nil
}
