package router

import (
	"final-project/controller"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	// swagger embed files
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	orderRouter := r.Group("/todos")
	{
		orderRouter.POST("/", controller.CreateTodo)
		orderRouter.GET("/", controller.GetTodos)
		orderRouter.GET("/:id", controller.GetTodo)
		orderRouter.PUT("/:id", controller.UpdateTodo)
		orderRouter.DELETE("/:id", controller.DeleteTodo)
	}
	return r
}
