// Package components
// @File    : pagehandler.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:03
// @Desc    :
package components

import (
	"asense/common/utils/characterutil"
	"strings"
)

type PageResult struct {
	Page     int     // 页码
	PageSize int     // 每页数量
	Filter   *string // 关键字
}

// PageHandle 分页处理
// page: 页码
// pageSize: 每页数量
// filter: 关键字
// return: 分页结果
func PageHandle(page, pageSize int, filter *string) PageResult {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize > 1000 {
		pageSize = 1000
	}

	if filter != nil {
		*filter = Filter(*filter)
	}

	return PageResult{
		Page:     (page - 1) * pageSize,
		PageSize: pageSize,
		Filter:   filter,
	}
}

// Filter 过滤关键字处理
// filter: 关键字
// return: 关键字结果
func Filter(filter string) string {
	_filter := strings.TrimSpace(filter)
	if len(_filter) > 0 {
		return characterutil.StitchingBuilderStr("%", _filter, "%")
	}
	return ""
}
