package requests

import (
	"bytes"
	"net/url"
	"sort"
	"strings"
)

// TopValues 自定义的请求参数类型
type TopValues map[string][]string

// Add 添加新的请求参数
func (vs TopValues) Add(k, v string) {
	vs[k] = append(vs[k], v)
}

// GetSortedKeys 获取已排序的KEY值
func (vs TopValues) GetSortedKeys() []string {
	return sortKeys(vs)
}

// GetUnsignedText 获取未签名的原始文本
func (vs TopValues) GetUnsignedText() string {
	sortedKeys := sortKeys(vs)
	buf := bytes.Buffer{}
	for _, k := range sortedKeys {
		buf.WriteString(k)
		buf.WriteString(strings.Join(vs[k], ","))
	}
	return buf.String()
}

// Encode 编码请求参数
func (vs TopValues) Encode() string {
	if vs == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(vs))
	for k := range vs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		vs := vs[k]
		keyEscaped := url.QueryEscape(k)
		buf.WriteString(keyEscaped + "=")

		comma := false
		for _, v := range vs {
			if comma {
				buf.WriteByte(',')
			}
			buf.WriteString(url.QueryEscape(v))
			comma = true
		}
	}
	return buf.String()
}

// ValidateRequired 校验必填项
func (vs TopValues) ValidateRequired(k string) {
	if v, ok := vs[k]; !ok || len(v) == 0 {
		panic("field: " + k + " is required.")
	}
}

// sortKeys 内部排序
func sortKeys(vs TopValues) []string {
	sortedKeys := make([]string, 0, len(vs))
	for k := range vs {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}
