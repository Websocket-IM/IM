package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, code int, message string, data interface{}) {

	fmt.Println(data, 8888)
	if err, ok := data.(error); ok {
		data = err.Error()
	}
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
