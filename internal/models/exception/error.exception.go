package exception

type ApiError struct {
	ErrorResponse ErrorResponse `json:"error"`
}

func (e *ApiError) Error() string {
	return e.ErrorResponse.Message
}

type ErrorResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Details *[]ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Reason    string   `json:"reason"`
	Reference []string `json:"reference"`
}
