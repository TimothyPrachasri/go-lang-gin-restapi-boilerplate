package main

import (
	"app/lib/interfaces/routes/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.ApplyRoutes(router)
	router.Run()
}
