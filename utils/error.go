package utils

type ApplicationError struct {
	ErrorCode       string
	Message         string
}

func SetError(errorCode string, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) Code() string {
	return e.ErrorCode
}


func GetErrorCode(err error) string {
	// Don't use FromError to avoid allocation of OK status.
	if err == nil {
		return ""
	}

	if se, ok := err.(interface {
		Code() string
	}); ok {
		return se.Code()
	}
	return ""
}