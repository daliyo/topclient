package responses

type TaobaoHttpDnsGetResponse struct {
	ErrorResponse

	Result string `json:"result,omitempty"`
}
