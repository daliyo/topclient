package responses

type TaobaoTimeGetResponse struct {
	ErrorResponse
	Time string `json:"time,omitempty"`
}
