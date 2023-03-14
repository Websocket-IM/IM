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
	user, _ := service.FindBy("username", rep.Username)
	fmt.Println(user, 9999)
	// 这里踩坑了,find方法返回的是已经初始化的切片，就算查不到数据也是空切片，而这个空切片不为Nil
	if len(user) > 0 {
		utils.JSON(c, 403, "error!", "该用户名已经被注册")
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

// 用户登录
func LoginUser(c *gin.Context) {
	user := model.LoginUserRep{}
	if err := utils.DefaultGetValidParams(c, &user); err != nil {
		utils.JSON(c, 400, "error!", err)
		return
	}
	// 用户名不存在错误
	if user, _ := service.FindBy("username", user.Username); len(user) == 0 {
		utils.JSON(c, 400, "error!", "用户名不存在")
		return
	}
	// 返回错误
	if err := service.LoginUser(&user); err != nil {
		utils.JSON(c, 404, "error!", err)
		return
	}
	utils.JSON(c, 200, "success!", "登录成功")
}

// 通过手机号登录
func LoginByPhoneCode(c *gin.Context) {
	phone := model.LoginByPhone{}
	if err := utils.DefaultGetValidParams(c, &phone); err != nil {
		utils.JSON(c, 400, "error!", err)
		return
	}
	utils.SMS(phone.Phone)
	utils.JSON(c, 200, "success!", "短信发送成功")

}

// 手机号验证码登录
func LoginByPhone(c *gin.Context) {
	loginByphonecode := model.LoginByPhoneCode{}
	if err := utils.DefaultGetValidParams(c, &loginByphonecode); err != nil {
		utils.JSON(c, 400, "error!", err)
		return
	}
	user, err := service.LoginByPhoneCode(&loginByphonecode)
	if err != nil {
		utils.JSON(c, 404, "error!", err)
		fmt.Println(err)
		return
	}
	accessTokenString, refreshTokenString := utils.GetToken(user, utils.RandNumber(10))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
			"user":         user,
		},
	})
}

// 根据字段查找
func FindBy(c *gin.Context) {

}
