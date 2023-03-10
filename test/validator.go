package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
	locale        = "chinese"
)

func TransInit(c *gin.Context) {
	//设置支持语言
	chinese := zh.New()
	english := en.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese, english)
	//设置验证器
	val := validator.New()
	//根据参数取翻译器实例
	trans, _ := uni.GetTranslator(locale)
	//翻译器注册到validator
	switch locale {
	case "chinese":
		zhTranslations.RegisterDefaultTranslations(val, trans)
		//使用fld.Tag.Get("comment")注册一个获取tag的自定义方法
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
	case "english":
		enTranslations.RegisterDefaultTranslations(val, trans)
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en_comment")
		})
	}
	c.Set(TranslatorKey, trans)
	c.Set(ValidatorKey, val)
}

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	c.ShouldBind(params)
	//获取验证器
	val, _ := c.Get(ValidatorKey)
	valid, _ := val.(*validator.Validate)
	//获取翻译器
	tran, _ := c.Get(TranslatorKey)
	trans, _ := tran.(ut.Translator)
	err := valid.Struct(params)
	//如果数据效验不通过，则将所有err以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
func loginTest(c *gin.Context) {
	TransInit(c)
	req := &login{}
	err := DefaultGetValidParams(c, req)
	//失败时输出
	if err != nil {
		c.JSON(404, gin.H{
			"code": 2000,
			"err":  err.Error(),
		})
		return
	}
	//成功时输出
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"name": req.Username,
	})
}

type login struct {
	Username string ` form:"username" comment:"用户名" validate:"required,min=2" `
}

func main() {
	route := gin.Default()
	route.POST("/login", loginTest)
	route.Run(":8099")
}
