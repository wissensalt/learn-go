package router

import (
	"os"

	"../controller"
	"../global"
	"../util"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRouter() {
	r := gin.Default()
	var port string
	err := godotenv.Load()

	if err != nil {
		global.Logger.Error("Error Open File Property for Server")
	} else {
		port = os.Getenv("server.port")
		sessionStore := sessions.NewCookieStore([]byte(os.Getenv("session.secret")))
		r.Use(sessions.Sessions(os.Getenv("session.name"), sessionStore))
		auth := r.Group("v1/auth")
		{
			auth.GET("/login", controller.Login)
			auth.GET("/findAll", controller.FindAll)
			auth.POST("/findById", controller.FindById)
			auth.POST("/insert", controller.InsertEmployee)
			auth.POST("/update", controller.UpdateEmployee)
			auth.POST("/delete", controller.DeleteEmployee)
		}

	}
	global.Logger.Info("#ROUTER INITATED")
	util.RollingLog()
	r.Run(port)
}
