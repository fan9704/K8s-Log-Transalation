package config

import (
	"k8s-log-translation/app/controller"

	"github.com/gin-gonic/gin"
)


func RouteUsers(r *gin.Engine) {
	posts := r.Group("/v1/users")
	{
		posts.POST("/", controller.NewUsersController().CreateUser)
		posts.GET("/:id", controller.QueryUsersController().GetUser)
	}
}
func RouteCommand(r *gin.Engine) {
	commands := r.Group("/v1/command")
	{
		commands.GET("/k8sgpt", controller.ExecuteCommandController().K8sGPTCommand)

	}
}