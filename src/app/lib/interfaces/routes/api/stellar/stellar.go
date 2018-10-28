package api

import (
	stellarController "app/lib/interfaces/controllers/stellar"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	stellar := r.Group("/stellar")
	{
		stellar.PUT("/transfer", stellarController.Transfer)
	}
}
