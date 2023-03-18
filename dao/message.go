package dao

import (
	"context"
	"fmt"
	"ginchat/common"
)

const (
	PublishKey = "websocket"
)

// Publish 发布消息到commom.RDBis
func Publish(ctx context.Context, channel string, msg string) error {
	fmt.Println("Publish：", msg)
	err := common.RDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅commom.RDBis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := common.RDB.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("lll")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe：", msg.Payload)
	return msg.Payload, err
}
