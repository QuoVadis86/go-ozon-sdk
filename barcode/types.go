package barcode

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"`
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}
