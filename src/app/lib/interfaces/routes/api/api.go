package api

import (
	stellarApi "app/lib/interfaces/routes/api/stellar"
	tokenApi "app/lib/interfaces/routes/api/token"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes take gin engine to create route named api and send it to each token api routes.
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		tokenApi.ApplyRoutes(api)
		stellarApi.ApplyRoutes(api)
	}
}
