package api

import (
	stellarController "app/lib/interfaces/controllers/stellar"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes apply routerGroup and create API endpoints based on existing group.
func ApplyRoutes(r *gin.RouterGroup) {
	stellar := r.Group("/stellar")
	{
		stellar.PUT("/transfer", stellarController.Transfer)
	}
}
