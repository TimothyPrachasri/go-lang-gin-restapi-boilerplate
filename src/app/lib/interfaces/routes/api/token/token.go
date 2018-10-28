package api

import (
	tokenController "app/lib/interfaces/controllers/token"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes apply routerGroup and create API endpoints based on existing group.
func ApplyRoutes(r *gin.RouterGroup) {
	token := r.Group("/token")
	{
		token.GET("/compare", tokenController.Compare)
	}
}
