package models

type SuccessResponse struct {
	Status  string `json:"status" default:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}
