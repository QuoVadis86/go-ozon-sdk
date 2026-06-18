package barcode

type V1AddBarcodeResult struct {
	Code string `json:"code"` // 错误代码。
	Error string `json:"error"` // 错误描述。
	Barcode string `json:"barcode"` // 未能绑定的条形码。
	SKU int64 `json:"sku"` // 未能绑定条形码的商品标识符。
}

type V1GenerateBarcodeRequest struct {
	ProductIds []string `json:"product_ids"` // 需要生成条形码的商品标识符。
}

type V1GenerateBarcodeResult struct {
	Code string `json:"code"` // 错误代码。
	Error string `json:"error"` // 错误描述。
	Barcode string `json:"barcode"` // 生成条形码时出错的条形码。
	ProductID int64 `json:"product_id"` // 未能成功生成条形码的商品标识符。
}

type V1GenerateBarcodeResponse struct {
	Errors []V1GenerateBarcodeResult `json:"errors"` // 生成条形码时出现的错误。
}

type V1Barcode struct {
	Barcode string `json:"barcode"` // 条形码的数值。不超过 100 个字符。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符 — SKU。
}

type V1AddBarcodeRequest struct {
	Barcodes []V1Barcode `json:"barcodes"` // 条形码与商品的列表。
}

type V1AddBarcodeResponse struct {
	Errors []V1AddBarcodeResult `json:"errors"` // 错误列表。
}
