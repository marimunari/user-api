package models

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	APIResponse
}

type SuccessResponse struct {
	APIResponse
}
