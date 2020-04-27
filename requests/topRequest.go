package requests

import (
	"net/http"
)

// ITopRequest TopRequest的接口定义
type ITopRequest interface {
	APIName() string
	GetForm() TopValues
	GetHeader() http.Header
	Validate()
}

// TopRequest 请求的基类
type TopRequest struct {
	Form   TopValues   //请求参数
	Header http.Header //请求头
}

func (req *TopRequest) GetForm() TopValues {
	return req.Form
}

func (req *TopRequest) GetHeader() http.Header {
	return req.Header
}
