package delivery

// 仓库位置。
type PolygonBindRequestwhLocation struct {
	Lon string `json:"lon"` // 仓库位置的地理经度。
	Lat string `json:"lat"` // 仓库位置的地理纬度。
}

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"` // 设施识别号。
	Time      int64 `json:"time"`       // 商品在该点快递到达的时间，以分钟计。
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID  int32                        `json:"delivery_method_id"` // 快递方法识别号。
	Polygons          []PolygonBindRequestpolygon  `json:"polygons"`           // 设施清单。
	WarehouseLocation PolygonBindRequestwhLocation `json:"warehouse_location"`
}

type Polygonv1Empty any

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"` // 快递设施的坐标，格式为 `[[[lat long]]]`。
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"` // 设施识别号。
}
