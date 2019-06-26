package util

import (
	"net/http"

	"../dto"
	"github.com/gin-gonic/gin"
)

func BuildResponseData(p_Context *gin.Context, p_ResponseData dto.ResponseData) {
	var result dto.ResponseData
	result.ResponseCode = p_ResponseData.ResponseCode
	result.ResponseMsg = p_ResponseData.ResponseMsg
	p_Context.JSON(http.StatusOK, result)
}

func BuildResponseDataWithContent(p_Context *gin.Context, p_ResponseData dto.ResponseData, p_Payload interface{}) {
	var result dto.ResponseDataContent
	result.ResponseData = p_ResponseData
	result.Content = p_Payload
	p_Context.JSON(http.StatusOK, result)
}
