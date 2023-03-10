package api

import (
	"ginchat/model"
	"ginchat/service"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
)

// 展示用户列表
func GetUserList(c *gin.Context) {

	if users, err := service.GetUserList(); err == nil {
		utils.JSON(c, 200, "用户已经注册", users)
	} else {
		utils.JSON(c, 500, "error!", err)
		return
	}
}

// 新增用户
func CreateUser(c *gin.Context) {
	// 迁移 schema

	rep := model.RegisterUserReq{}

	if err := utils.DefaultGetValidParams(c, &rep); err != nil {

		utils.JSON(c, 404, "error!", err)
		return
	}

	if err := service.CreateUser(&rep); err != nil {

		utils.JSON(c, 404, "error!", err)
		return
	}
	utils.JSON(c, 200, "success!", rep)

}
