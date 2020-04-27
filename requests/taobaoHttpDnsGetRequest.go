package requests

import (
	"net/http"
)

// TaobaoHttpDnsGetRequest 获取TOP DNS配置
type TaobaoHttpDnsGetRequest struct {
	TopRequest
}

// NewTaobaoTimeGetRequest 创建获取淘宝系统当前时间请求的新实例
func NewTaobaoHttpDnsGetRequest() *TaobaoHttpDnsGetRequest {
	req := &TaobaoHttpDnsGetRequest{}
	req.Form = make(TopValues)
	req.Header = make(http.Header)
	return req
}

// APIName 获取接口的名称
func (req *TaobaoHttpDnsGetRequest) APIName() string {
	return "taobao.httpdns.get"
}

func (req *TaobaoHttpDnsGetRequest) Validate() {

}
