package types

type V1Empty interface{}

type RpcStatus struct {
	Details []interface{} `json:"details"`
	Message string `json:"message"`
	Code int32 `json:"code"`
}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type ProtobufAny struct {
	TypeUrl string `json:"typeUrl"`
	Value string `json:"value"`
}
