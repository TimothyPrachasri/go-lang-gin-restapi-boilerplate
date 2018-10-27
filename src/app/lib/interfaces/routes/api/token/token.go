package token

import (
	"github.com/gin-gonic/gin"

	"app/lib/interfaces/controllers"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	token := r.Group("/token")
	{
		token.GET("/compare", tokencontroller.Compare)
	}
}
