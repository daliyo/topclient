package responses

type ErrorResponse struct {
	Error ErrorObject `json:"error_response,omitempty"`
}

type ErrorObject struct {
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	SubCode   string `json:"sub_code,omitempty"`
	SubMsg    string `json:"sub_msg,omitempty"`
}
