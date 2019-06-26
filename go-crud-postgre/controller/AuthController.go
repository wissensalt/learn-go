package controller

import (
	"../dto"
	"../util"
	"github.com/gin-gonic/gin"
)

func Login(p_Context *gin.Context) {
	ControllerLogger.Info("Login")

	var responseData dto.ResponseData
	responseData.ResponseCode = "200"
	responseData.ResponseMsg = "Oke"
	util.BuildResponseData(p_Context, responseData)
}
