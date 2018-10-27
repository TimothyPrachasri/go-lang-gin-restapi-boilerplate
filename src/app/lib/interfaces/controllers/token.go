package tokencontroller

import (
	"github.com/gin-gonic/gin"
)

func Compare(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
