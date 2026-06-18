package delivery

type Polygonv1Empty interface{}

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"`
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"`
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"`
	Polygons []interface{} `json:"polygons"`
	WarehouseLocation interface{} `json:"warehouse_location"`
}
