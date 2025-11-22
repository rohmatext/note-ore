package validator

type ValidationError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
