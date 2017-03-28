package httperror

type HttpError struct {
	Message    string     `json:"message"`
	Code       string     `json:"code"`
	Context    string     `json:"context"`
	StatusCode int        `json:"-"`
	Stack      StackTrace `json:"-"`
}

func NewHttpError(message string, code string, httpStatus int) *HttpError {
	return &HttpError{
		Message:    message,
		Code:       code,
		StatusCode: httpStatus,
	}
}

func (e HttpError) Error() string {
	return e.Message
}
