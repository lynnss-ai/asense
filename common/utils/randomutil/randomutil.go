// Package randomutil
// @File    : randomutil.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 14:47
// @Desc    :
package randomutil

import (
	"math/rand"
	"time"
)

// GetRandomNumStr 随机生成指定长度的数字字符串
// length: 需要生成的字符串长度
// return: 生成的字符串
func GetRandomNumStr(length int) string {
	const letterBytes = "0123456789"
	b := make([]byte, length)
	rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

// GetRandomStr 随机生成指定长度的字符串
// length: 需要生成的字符串长度
// return: 生成的字符串
func GetRandomStr(length int) string {
	const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
