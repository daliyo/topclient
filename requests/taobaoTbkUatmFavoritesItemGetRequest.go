package requests

import (
	"net/http"
	"strconv"
)

type TaobaoTbkUatmFavoritesItemGetRequest struct {
	TopRequest
}

func NewTaobaoTbkUatmFavoritesItemGetRequest() *TaobaoTbkUatmFavoritesItemGetRequest {
	req := &TaobaoTbkUatmFavoritesItemGetRequest{}
	req.Form = make(TopValues)
	req.Header = make(http.Header)

	req.Form.Add("fields", "title,pict_url,zk_final_price,click_url,volume,category,coupon_click_url,coupon_info,")

	return req
}

// APIName 获取接口的名称
func (req *TaobaoTbkUatmFavoritesItemGetRequest) APIName() string {
	return "taobao.tbk.uatm.favorites.item.get"
}

func (req *TaobaoTbkUatmFavoritesItemGetRequest) Validate() {
	req.Form.ValidateRequired("adzone_id")
	req.Form.ValidateRequired("favorites_id")
}

func (req *TaobaoTbkUatmFavoritesItemGetRequest) SetAdzoneID(adzoneId int) {
	req.Form.Add("adzone_id", strconv.Itoa(adzoneId))
}

func (req *TaobaoTbkUatmFavoritesItemGetRequest) SetFavoritesID(favoritesId int) {
	req.Form.Add("favorites_id", strconv.Itoa(favoritesId))
}

// SetPageNo 设置页码
func (req *TaobaoTbkUatmFavoritesItemGetRequest) SetPageNo(page int) {
	req.Form.Add("page_no", strconv.Itoa(page))
}

// SetPageSize 设置页大小
func (req *TaobaoTbkUatmFavoritesItemGetRequest) SetPageSize(size int) {
	req.Form.Add("page_size", strconv.Itoa(size))
}
