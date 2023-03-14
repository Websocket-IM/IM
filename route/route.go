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
		user.POST("/login", api.LoginUser)
		user.POST("/login/phone/code", api.LoginByPhoneCode) //发验证码
		user.POST("/login/phone", api.LoginByPhone)          //校对验证
		user.POST("/find", api.FindBy)
	}

	return r
}
