package controller

import (
	"../global"
	"../util"
	"github.com/gin-gonic/gin"
)

func Login(p_Context *gin.Context) {
	global.Logger.Info("Login")
	util.BuildResponseOK(p_Context, "200", "OK")
}
