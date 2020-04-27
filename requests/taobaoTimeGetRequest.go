package requests

// TaobaoTimeGetRequest 获取淘宝系统当前时间
type TaobaoTimeGetRequest struct {
	TopRequest
}

// NewTaobaoTimeGetRequest 创建获取淘宝系统当前时间请求的新实例
func NewTaobaoTimeGetRequest() *TaobaoTimeGetRequest {
	req := &TaobaoTimeGetRequest{}
	req.Form = make(TopValues)
	return req
}

// APIName 获取接口的名称
func (req *TaobaoTimeGetRequest) APIName() string {
	return "taobao.time.get"
}

func (req *TaobaoTimeGetRequest) Validate() {

}
