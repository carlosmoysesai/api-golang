package routes

import (
	"github.com/carlosmoysesai/api-golang/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
	router.GET("/:id", controller.GetUser)
}
