// Package characterutil
// @File    : characterutil.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 14:44
// @Desc    :
package characterutil

import (
	"encoding/json"
	"strings"
)

// StitchingBuilderStr 字符串拼接
// args: 需要拼接的字符串
// return: 拼接后的字符串
func StitchingBuilderStr(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}

// IsJSON 判断是否是json格式
// data: 需要判断的数据
// return: 是否是json格式
func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}
