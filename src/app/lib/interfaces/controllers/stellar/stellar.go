package controller

import (
	"net/http"

	service "app/lib/domain/services/stellar"

	"github.com/gin-gonic/gin"
)

// Transfer is a function used for transfering amount of lumens from source seed to destination address.
// if the process of transfer has been successfully proceeded the respond will be status true else will be err.
func Transfer(c *gin.Context) {
	from := c.PostForm("from")
	to := c.PostForm("to")
	amount := c.PostForm("amount")
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
