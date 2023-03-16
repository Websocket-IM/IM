package api

import (
	"fmt"
	"ginchat/external"
	"ginchat/model"
	"ginchat/service"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var conf model.GithubConf

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
	//phone := model.LoginByPhone{}
	//
	//if err := utils.DefaultGetValidParams(c, &phone); err != nil {
	//	utils.JSON(c, 400, "error!", err)
	//	return
	//}
	phone := c.Query("phone")
	fmt.Println(phone, "电话号码")
	external.SMS(phone)
	utils.JSON(c, 200, "success!", "短信发送成功")

}

// 手机号验证码登录
func LoginByPhone(c *gin.Context) {
	loginByphonecode := model.LoginByPhoneCode{}
	if err := utils.DefaultGetValidParams(c, &loginByphonecode); err != nil {
		utils.JSON(c, 400, "error!", err)
		fmt.Println(1111)
		return
	}
	fmt.Println(loginByphonecode)
	user, err := service.LoginByPhoneCode(&loginByphonecode)
	if err != nil {
		utils.JSON(c, 404, "error!", err)
		fmt.Println(222222)
		return
	}
	accessTokenString, refreshTokenString := utils.GetToken(user.ID, utils.RandNumber(10))
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

// 处理第三方 Github 登录请求
func HandleGithubLogin(c *gin.Context) {
	if err := viper.UnmarshalKey("github", &conf); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %s", err))
	}
	authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", conf.ClientId, conf.RedirectUrl)
	c.Redirect(307, authURL)

}

// 处理 GitHub 登录回调请求
func HandleGithubCallback(c *gin.Context) {
	// 从查询参数中获取授权码
	code := c.Query("code")

	// 交换授权码获取访问令牌
	tokenAuthUrl := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code)
	// 获取 token
	var token *model.GithubToken
	var err error
	if token, err = external.GetGithubToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("------%v", token)
	// 通过token，获取用户信息
	var user model.User
	if user, err = external.GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}
	if err = service.AddUser(&user); err != nil {
		fmt.Println(err)
		utils.JSON(c, 404, "error", "github新增用户失败")
		return
	}
	// 返回token和用户信息
	accessTokenString, refreshTokenString := utils.GetToken(user.ID, utils.RandNumber(10))
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
