package topsdk

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/daliyo/topclient/constants"
	"github.com/daliyo/topclient/requests"
)

var (
	signMethod     = constants.SIGN_METHOD_HMAC
	responseFormat = constants.FORMAT_JSON
	version        = "2.0"
	simplify       = true
)

// TopClient 淘宝开放平台客户端
type TopClient struct {
	apiGateway string
	appKey     string
	appSecret  string
}

func verifyRequired(m map[string]string) error {
	for k, v := range m {
		if v == "" {
			return errors.New(k + " is required")
		}
	}
	return nil
}

// NewTopClient 创建 TopClient 的新实例
func NewTopClient(apiGateway, appKey, appSecret string) (*TopClient, error) {

	err := verifyRequired(map[string]string{
		"apiGateway": apiGateway,
		"appKey":     appKey,
		"appSecret":  appSecret,
	})

	if err != nil {
		return nil, err
	}

	return &TopClient{
		apiGateway, appKey, appSecret,
	}, nil
}

// Do 执行请求
func (c TopClient) Do(r requests.ITopRequest, respModel interface{}) error {

	r.Validate() //验证参数有效性

	form := r.GetForm()

	err := signForm(form, r.APIName(), c.appKey, c.appSecret) //对Form签名
	if err != nil {
		panic(err)
	}

	requestURL := makeRequestURL(c.apiGateway, form)
	req, err := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	if header := r.GetHeader(); header != nil {
		req.Header = header
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return parseResponseBody(resp, &respModel)
}

// parseResponseBody 解析响应结果的正文内容
func parseResponseBody(resp *http.Response, respModel interface{}) error {

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch responseFormat {
	case constants.FORMAT_JSON:
		return json.Unmarshal(buf, &respModel)
	case constants.FORMAT_XML:
		panic("unimplemented XML parser.")
	default:
		panic("unimplemented " + responseFormat + " parser.")
	}
}

// makeRequestURL 构造请求地址
func makeRequestURL(apiGateway string, form requests.TopValues) string {
	return fmt.Sprintf("%v?%v", apiGateway, form.Encode())
}

// signForm 签名请求参数
func signForm(form requests.TopValues, methodName, appKey, appSecret string) error {

	//填充公共请求参数
	form.Add(constants.METHOD, methodName)
	form.Add(constants.APP_KEY, appKey)
	form.Add(constants.SIGN_METHOD, signMethod)
	form.Add(constants.TIMESTAMP, time.Now().Format(constants.TIMESTAMP_FORMAT))
	form.Add(constants.FORMAT, responseFormat)
	form.Add(constants.VERSION, version)
	form.Add(constants.SIMPLIFY, strconv.FormatBool(simplify))

	sign, err := sumHash(form, appSecret)
	if err != nil {
		return err
	}
	form.Add(constants.SIGN, sign)
	return nil
}

// sumHash 计算哈希值
func sumHash(form requests.TopValues, secret string) (string, error) {
	secretBuf := []byte(secret)
	unsignedText := form.GetUnsignedText()

	var h hash.Hash
	switch signMethod {
	case constants.SIGN_METHOD_HMAC:
		h = hmac.New(md5.New, secretBuf)
		_, err := h.Write([]byte(unsignedText))
		if err != nil {
			return "", err
		}

	case constants.SIGN_METHOD_MD5:
		h = md5.New()
		if _, err := h.Write([]byte(secret)); err != nil {
			return "", err
		}
		if _, err := h.Write([]byte(unsignedText)); err != nil {
			return "", err
		}
		if _, err := h.Write([]byte(secret)); err != nil {
			return "", err
		}

	default:
		panic("unsupported hash algorithm")
	}

	sum := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(sum)), nil
}
