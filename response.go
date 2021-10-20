package response

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ApiError) Error() string {
	return e.Msg
}

type errorWrapper struct {
	ApiError
	Data interface{} `json:"data"`
}

func Render(err ApiError, args ...interface{}) interface{} {
	if len(args) == 1 {
		return errorWrapper{ApiError: err, Data: args[0]}
	}
	return errorWrapper{ApiError: err, Data: args}
}
