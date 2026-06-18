package delivery

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"` // 快递设施的坐标，格式为 `[[[lat long]]]`。
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"` // 设施识别号。
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"` // 快递方法识别号。
	Polygons []interface{} `json:"polygons"` // 设施清单。
	WarehouseLocation interface{} `json:"warehouse_location"`
}

type Polygonv1Empty interface{}
