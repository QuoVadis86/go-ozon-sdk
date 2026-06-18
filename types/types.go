package types

type V1Empty interface{}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type ProtobufAny struct {
	TypeUrl string `json:"typeUrl"` // 数据传输协议类型。
	Value string `json:"value"` // 错误的值。
}

type RpcStatus struct {
	Code int32 `json:"code"` // 错误代码。
	Details []interface{} `json:"details"` // 有关错误的补充信息。
	Message string `json:"message"` // 错误描述。
}
