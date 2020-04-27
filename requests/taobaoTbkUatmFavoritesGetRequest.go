package requests

import (
	"net/http"
)

type TaobaoTbkUatmFavoritesGetRequest struct {
	TopRequest
}

func NewTaobaoTbkUatmFavoritesGetRequest() *TaobaoTbkUatmFavoritesGetRequest {
	req := &TaobaoTbkUatmFavoritesGetRequest{}
	req.Form = make(TopValues)
	req.Header = make(http.Header)

	req.Form.Add("page_no", "1")
	req.Form.Add("page_size", "30")
	//req.Form.Add("type", "2") //仅高佣选品组
	req.Form.Add("fields", "favorites_title,favorites_id,type")
	return req
}

// APIName 获取接口的名称
func (req *TaobaoTbkUatmFavoritesGetRequest) APIName() string {
	return "taobao.tbk.uatm.favorites.get"
}

func (req *TaobaoTbkUatmFavoritesGetRequest) Validate() {
}
