package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("employee", controller.FindAll)
		v1.POST("employee", controller.Insert)
		v1.GET("employee/:id", controller.FindById)
		v1.PUT("employee/:id", controller.Update)
		v1.DELETE("employee/:id", controller.Delete)
	}

	return r
}
