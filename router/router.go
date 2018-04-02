package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/controller"
)

func Initialize(r *gin.Engine) {
	r.GET("/", controller.APIEndpoints)

	//api := r.Group("")
	{

		//api.GET("/service", controller.GetEmails)
		//api.GET("/emails/:id", controllers.GetEmail)
		//api.POST("/emails", controllers.CreateEmail)
		//api.PUT("/emails/:id", controllers.UpdateEmail)
		//api.DELETE("/emails/:id", controllers.DeleteEmail)
		//
		//api.GET("/users", controllers.GetUsers)
		//api.GET("/users/:id", controllers.GetUser)
		//api.POST("/users", controllers.CreateUser)
		//api.PUT("/users/:id", controllers.UpdateUser)
		//api.DELETE("/users/:id", controllers.DeleteUser)

	}
}