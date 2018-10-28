package controller

import (
	"fmt"
	"net/http"

	service "app/lib/domain/services/stellar"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	from := c.PostForm("from")
	to := c.PostForm("to")
	amount := c.PostForm("amount")
	fmt.Println(from, to, amount, "test")
	isValid := from != "" || to != "" || amount != ""
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Insufficient Form Request",
		})
	}
	status, err := service.Stellar{}.SendLumens(amount, from, to)
	if status {
		c.JSON(http.StatusOK, gin.H{
			"status": status,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err,
	})
}
