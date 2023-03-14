package utils

import (
	"errors"
	"fmt"
	"ginchat/model"
	"ginchat/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var accessSecret = []byte("liuxian")
var refreshSecret = []byte("123")

// GetToken 获取accessToken和refreshToken
func GetToken(user model.User, state string) (string, string) {
	// accessToken 的数据
	aT := model.MyClaims{
		user,
		state,
		jwt.StandardClaims{
			Issuer:    "AR",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
	}
	// refreshToken 的数据
	rT := model.MyClaims{
		user,
		state,
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

func ParseToken(accessTokenString, refreshTokenString string) (*model.MyClaims, bool, error) {
	fmt.Println("In ParseToken")
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if claims, ok := accessToken.Claims.(*model.MyClaims); ok && accessToken.Valid {
		return claims, false, nil
	}

	fmt.Println("RefreshToken")
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims, ok := refreshToken.Claims.(*model.MyClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}

	return nil, false, errors.New("invalid token")
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
			parts[1], parts[2] = GetToken(parseToken.User, parseToken.State)
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
		users, err := service.FindBy("phone", parseToken.Phone)
		if err != nil {
			JSON(c, 404, "error!", "重新赋值token里面的user出错")
		}
		parseToken.User = users[0]
		c.Set("user", parseToken.User)
		c.Next()
	}
}
