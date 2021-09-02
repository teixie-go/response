package response

var (
	StatusOK = 200
	MsgOK    = "success"
)

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

func ErrorMessage(err ApiError, msg string, args ...interface{}) interface{} {
	if msg != "" {
		err.Msg = msg
	}
	if len(args) == 1 {
		return errorWrapper{err, args[0]}
	}
	return errorWrapper{err, args}
}

func Error(err ApiError, args ...interface{}) interface{} {
	if len(args) == 1 {
		return errorWrapper{err, args[0]}
	}
	return errorWrapper{err, args}
}

func Ok(args ...interface{}) interface{} {
	return Error(ApiError{StatusOK, MsgOK}, args...)
}
