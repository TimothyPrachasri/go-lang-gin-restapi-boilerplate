package main

import (
	api "app/lib/interfaces/routes/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// apply router ability to api
	api.ApplyRoutes(router)
	router.Run()
}
