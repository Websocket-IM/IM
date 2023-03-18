package api

import (
	"fmt"
	"ginchat/common"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 发送消息
func SendMessage(c *gin.Context) {
	fmt.Println(1111)
	upGrande := websocket.Upgrader{
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		//设置请求协议
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}
	//创建连接
	ws, err := upGrande.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		common.SugarLogger.Error("websocket连接失败.err : %v", err)
		return
	}
	fmt.Println("连接成功")
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	service.MessageHandler(c, ws)
}
