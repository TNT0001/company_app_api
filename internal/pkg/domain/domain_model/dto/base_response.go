package dto

// BaseResponse struct
type BaseResponse struct {
	Status int            `json:"status"`
	Result interface{}    `json:"result"`
	Error  *ErrorResponse `json:"error"`
}

// ErrorResponse struct
type ErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
