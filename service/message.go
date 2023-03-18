package service

import (
	"fmt"
	"ginchat/common"
	"ginchat/dao"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

func MessageHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := dao.Subscribe(c, dao.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
			common.SugarLogger.Error(err)
		}
	}
}
