package delivery

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"` // 设施识别号。
}

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"` // 设施识别号。
	Time      int64 `json:"time"`       // 商品在该点快递到达的时间，以分钟计。
}

// 仓库位置。
type PolygonBindRequestwhLocation struct {
	Lat string `json:"lat"` // 仓库位置的地理纬度。
	Lon string `json:"lon"` // 仓库位置的地理经度。
}

type Polygonv1PolygonBindRequest struct {
	Polygons          []PolygonBindRequestpolygon  `json:"polygons"` // 设施清单。
	WarehouseLocation PolygonBindRequestwhLocation `json:"warehouse_location"`
	DeliveryMethodID  int32                        `json:"delivery_method_id"` // 快递方法识别号。
}

type Polygonv1Empty any

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"` // 快递设施的坐标，格式为 `[[[lat long]]]`。
}
