package api

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context) {
	c.JSON(200, "welcome")
}
