package api

import (
	tokenController "app/lib/interfaces/controllers/token"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	token := r.Group("/token")
	{
		token.GET("/compare", tokenController.Compare)
	}
}
