package requests

import (
	"net/http"
	"strconv"
)

type TaobaoTbkDgMaterialOptionalRequest struct {
	TopRequest
}

func NewTaobaoTbkDgMaterialOptionalRequest() *TaobaoTbkDgMaterialOptionalRequest {
	req := &TaobaoTbkDgMaterialOptionalRequest{}
	req.Form = make(TopValues)
	req.Header = make(http.Header)
	return req
}

// APIName 获取接口的名称
func (req *TaobaoTbkDgMaterialOptionalRequest) APIName() string {
	return "taobao.tbk.dg.material.optional"
}

func (req *TaobaoTbkDgMaterialOptionalRequest) Validate() {
	req.Form.ValidateRequired("adzone_id")
}

// SetQ 设置商品筛选-查询词
func (req *TaobaoTbkDgMaterialOptionalRequest) SetQ(q string) {
	req.Form.Add("q", q)
}

// SetCat 设置商品筛选-类目ID
func (req *TaobaoTbkDgMaterialOptionalRequest) SetCat(cat string) {
	req.Form.Add("cat", cat)
}

// SetAdzoneID 设置adzone_id mm_xxx_xxx_12345678三段式的最后一段数字
func (req *TaobaoTbkDgMaterialOptionalRequest) SetAdzoneID(adzoneID int) {
	req.Form.Add("adzone_id", strconv.Itoa(adzoneID))
}

// SetPageNo 设置页码
func (req *TaobaoTbkDgMaterialOptionalRequest) SetPageNo(page int) {
	req.Form.Add("page_no", strconv.Itoa(page))
}

// SetPageSize 设置页大小
func (req *TaobaoTbkDgMaterialOptionalRequest) SetPageSize(size int) {
	req.Form.Add("page_size", strconv.Itoa(size))
}
