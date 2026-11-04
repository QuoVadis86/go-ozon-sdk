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
	if len(decoded.SKU) == 0 {
		return nil
	}
	if decoded.SKU[0] == '"' {
		return json.Unmarshal(decoded.SKU, &result.SKU)
	}
	var sku json.Number
	if err := json.Unmarshal(decoded.SKU, &sku); err != nil {
		return err
	}
	result.SKU = sku.String()
	return nil
}
