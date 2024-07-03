package main

import (
	"github.com/carlosmoysesai/api-golang/config"
	"github.com/carlosmoysesai/api-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}
