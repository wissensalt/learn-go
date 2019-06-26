package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func InitV1() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/employee")
	{
		v1.POST("login", controller.Login)
		v1.POST("/visitor/get")
		v1.GET("/findall", controller.FindAll)
		v1.POST("/findbyid", controller.FindById)
		v1.POST("/insert", controller.Insert)
	}

	v2 := r.Group("/v1/employees")
	{
		v2.GET("/hello", controller.HelloWorld)
	}
	return r
}
