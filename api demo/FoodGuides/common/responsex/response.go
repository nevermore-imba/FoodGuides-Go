package responsex

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(resData interface{}, message string) Body {
	return Body{Code: 1, Message: message, Data: resData}
}

func FailureResponse(resData interface{}, message string, code int) Body {
	return Body{Code: code, Message: message, Data: resData}
}
