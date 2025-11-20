package validator

type ValidationError struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
