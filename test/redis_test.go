package test

import (
	"context"
	"fmt"
	"ginchat/dao"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// TestPublish 测试发布消息到redis
func TestPublish(t *testing.T) {
	msg := "当前时间: " + time.Now().Format("15:04:05")
	err := dao.Publish(ctx, dao.PublishKey, msg)
	if err != nil {

		fmt.Println(err)
	}
}
