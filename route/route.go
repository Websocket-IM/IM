package route

import (
	"ginchat/api"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.GET("/index", api.GetIndex)
	//用户模块
	user := r.Group("/user")
	{
		user.POST("/getUserList", api.GetUserList)
		user.POST("/createUser", api.CreateUser)
		user.DELETE("/deleteUser", api.DeleteUser)
		user.POST("/updateUser", api.UpadateUser)
	}

	return r
}
