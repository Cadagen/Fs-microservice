package utils

type CliResponse[T any] struct {
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
	Data    *T     `json:"data,omitempty"`
}

func NewCliResponse[T any](data T) CliResponse[T] {
	return CliResponse[T]{
		Type: "ok",
		Data: &data,
	}
}

func NewCliError(msg string) CliResponse[any] {
	return CliResponse[any]{
		Type:    "error",
		Message: msg,
	}
}
