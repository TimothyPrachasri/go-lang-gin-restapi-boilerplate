package api

import (
	stellarApi "app/lib/interfaces/routes/api/stellar"
	tokenApi "app/lib/interfaces/routes/api/token"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		tokenApi.ApplyRoutes(api)
		stellarApi.ApplyRoutes(api)
	}
}
