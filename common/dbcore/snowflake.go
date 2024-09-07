// Package dbcore
// @File    : snowflake.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 14:42
// @Desc    :
package dbcore

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"sync"
)

var (
	node  *snowflake.Node
	once  sync.Once
	mutex sync.Mutex
)

// NewId 使用雪花算法获取唯一ID标识
func NewId() string {
	once.Do(func() {
		for retry := 0; retry < 10; retry++ {
			n, err := snowflake.NewNode(int64(rand.Int() % 1024))
			if err != nil {
				continue
			}
			node = n
			break
		}
		if node == nil {
			panic("snowflake id generate error")
		}
	})

	mutex.Lock()
	defer mutex.Unlock()

	return string(node.Generate().Bytes())
}
