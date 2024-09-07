// Package timeutil
// @File    : timeutil.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 14:48
// @Desc    :
package timeutil

import "time"

// TimeFormat 将time.Time时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
// 如果时间为空，则返回"——"
func TimeFormat(t *time.Time) string {
	if t == nil {
		return "——"
	}
	return t.Format(time.DateTime)
}

// TimeMilliFormat 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
// 如果时间为空，则返回"——"
func TimeMilliFormat(t *int64) string {
	if t == nil || *t == 0 {
		return "——"
	}
	return time.Unix(*t/1e3, 0).Format(time.DateTime)
}

// TimeMilliFormatToYMD 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD"
// 如果时间为空，则返回"——"
func TimeMilliFormatToYMD(t *int64) string {
	if t == nil || *t == 0 {
		return "——"
	}
	return time.Unix(*t/1e3, 0).Format(time.DateOnly)
}
