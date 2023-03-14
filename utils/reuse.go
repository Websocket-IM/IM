package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"time"
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

func RandNickname() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100000000)
	return fmt.Sprintf("user%08d", num)
}

func RandNumber(n int) string {
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow10(n - 1))
	max := int(math.Pow10(n) - 1)
	return fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
}
