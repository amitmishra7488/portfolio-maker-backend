package utils

type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
    Message string            `json:"message"`
    Errors  map[string]string `json:"errors,omitempty"`
}
