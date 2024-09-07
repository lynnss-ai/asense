// Package errorx
// @File    : errorx.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 14:43
// @Desc    :
package errorx

import (
	"asense/common/utils/characterutil"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	defaultCode                   = 1001 //默认错误
	defaultGrpcCode               = 1002 //Grpc错误
	databaseErrCode               = 1003 //数据库查询错误
	databaseErrRecordNotFoundCode = 1004 //数据库查询数据为空错误
)

type CodeErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func NewDataBaseError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewCodeError(databaseErrRecordNotFoundCode, "查询[Database]数据为nil")
	} else {
		return NewCodeError(databaseErrCode, characterutil.StitchingBuilderStr("Database操作失败:", err.Error()))
	}
}

func NewGRPCError(err error) error {
	return NewCodeError(defaultGrpcCode, status.Convert(err).Message())
}

func NewCodeError(code int, msg string) error {
	return &CodeErr{
		Code: code,
		Msg:  msg,
	}
}

func (e *CodeErr) Error() string {
	return e.Msg
}

func (e *CodeErr) Data() *CodeErr {
	return &CodeErr{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
