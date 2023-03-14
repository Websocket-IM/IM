package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// User 用户信息
type User struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

type MyClaims struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	jwt.StandardClaims
}

var accessSecret = []byte("liuxian")
var refreshSecret = []byte("123")

// GetToken 获取accessToken和refreshToken
func GetToken(name, pwd string) (string, string) {
	// accessToken 的数据
	aT := MyClaims{
		name,
		pwd,
		jwt.StandardClaims{
			Issuer:    "AR",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
	}
	// refreshToken 的数据
	rT := MyClaims{
		name,
		pwd,
		jwt.StandardClaims{
			Issuer:    "AR",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aT)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rT)
	accessTokenSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", ""
	}
	refreshTokenSigned, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", ""
	}
	return accessTokenSigned, refreshTokenSigned
}

func ParseToken(accessTokenString, refreshTokenString string) (*MyClaims, bool, error) {
	fmt.Println("In ParseToken")
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if claims, ok := accessToken.Claims.(*MyClaims); ok && accessToken.Valid {
		return claims, false, nil
	}

	fmt.Println("RefreshToken")
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims, ok := refreshToken.Claims.(*MyClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}

	return nil, false, errors.New("invalid token")
}

func authHandler(c *gin.Context) {
	fmt.Println("In authHandler")
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 2001,
			"msg":  "无效参数",
		})
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user = ", user)
	if !(user.Name == "ar" && user.Pwd == "123456") {
		c.JSON(200, gin.H{
			"code": 2002,
			"msg":  "鉴权失败",
		})
		fmt.Println("User not exist or password error")
		return
	}
	accessTokenString, refreshTokenString := GetToken(user.Name, user.Pwd)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		},
	})
}

// JWTAuthMiddleware 用鉴权到中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 默认双Token放在请求头Authorization的Bearer中，并以空格隔开
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(c.Request.Header)
		if authHeader == "" {
			c.JSON(200, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		fmt.Println("authHeader = ", authHeader)
		parts := strings.Split(authHeader, " ")
		fmt.Println("len = ", len(parts))
		fmt.Println("parts[0] = ", parts[0])
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			c.JSON(200, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		parseToken, isUpd, err := ParseToken(parts[1], parts[2])
		if err != nil {
			c.JSON(200, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// accessToken 已经失效，需要刷新双Token
		if isUpd {
			parts[1], parts[2] = GetToken(parseToken.Name, parseToken.Pwd)
			// 如果需要刷新双Token时，返回双Token
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "鉴权成功",
				"data": gin.H{
					"accessToken":  parts[1],
					"refreshToken": parts[2],
				},
			})
		}
		c.Set("username", parseToken.Name)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	// 获取Token
	r.POST("/auth", authHandler)
	// Token访问测试
	r.GET("/home", JWTAuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"username": username},
		})
	})
	r.Run()
}
