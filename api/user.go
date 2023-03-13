package api

import (
	"fmt"
	"ginchat/model"
	"ginchat/service"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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

// 删除用户
func DeleteUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("id"))
	user := model.User{}
	user.ID = uint(uid)
	err := service.DeleteUser(&user)
	if err != nil {
		utils.JSON(c, 404, "error!", err)
		return
	}
	utils.JSON(c, 200, "success!", "删除用户成功")
}

// 更新用户资料
func UpadateUser(c *gin.Context) {
	user := model.UpadateUserRep{}
	if err := utils.DefaultGetValidParams(c, &user); err != nil {
		utils.JSON(c, 404, "error!", err)
		return
	}
	fmt.Println("upadatuser:   ", user)

	err := service.UpadateUser(&user)
	if err != nil {
		utils.JSON(c, 404, "error!", err)
		return
	}
	utils.JSON(c, 200, "success!", "更新用户资料成功")

}
