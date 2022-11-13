package response

const (
	CodeSuccess      = "00000"
	CodeNoPermission = "A0100"
	CodeParamError   = "A0200"
	CodeNotFound     = "A0400"
	CodeRateLimit    = "A0429"
	CodeSystemError  = "B0100"

	MessageSuccess      = "成功"
	MessageSystemError  = "系统执行出错"
	MessageNoPermission = "没有权限"
	MessageNotFound     = "接口不存在"
	MessageRateLimit    = "访问过于频繁"
	MessageParamError   = "参数错误"

	MessageDefault = "未知错误"
)

var (
	msgMap = map[string]string{
		CodeSuccess:      MessageSuccess,
		CodeSystemError:  MessageSystemError,
		CodeNoPermission: MessageNoPermission,
		CodeNotFound:     MessageNotFound,
		CodeParamError:   MessageParamError,
		CodeRateLimit:    MessageRateLimit,
	}
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func GetResponseMessage(code string) string {
	v, ok := msgMap[code]
	if ok {
		return v
	} else {
		return MessageDefault
	}
}

func SuccessResponse(result any) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: GetResponseMessage(CodeSuccess),
		Result:  result,
	}
}

func DefaultSuccessResponse() *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: GetResponseMessage(CodeSuccess),
		Result:  nil,
	}
}

func ErrorResponse(code string, result any) *Response {
	return &Response{
		Code:    code,
		Message: GetResponseMessage(code),
		Result:  result,
	}
}

func DefaultErrorResponse(code string) *Response {
	return &Response{
		Code:    code,
		Message: GetResponseMessage(code),
		Result:  nil,
	}
}
