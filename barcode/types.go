package barcode

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"` // 生成条形码时出现的错误。
}

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"` // 条形码与商品的列表。
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"` // 错误列表。
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"` // 需要生成条形码的商品标识符。
}
