package presenter

type ApiResponse[T any] struct {
	Message string    `json:"message"`
	Data    T         `json:"data"`
	Meta    *PageMeta `json:"meta,omitempty"`
}

type PageMeta struct {
	Cursor *string `json:"cursor"`
}
