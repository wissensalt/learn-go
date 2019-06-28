package util

import (
	"net/http"

	"../dto"
	"github.com/gin-gonic/gin"
)

func BuildResponseOK(p_Context *gin.Context, p_ResponseCode string, p_ResponseMsg string) {
	p_Context.JSON(http.StatusOK, BuildResponseData(p_ResponseCode, p_ResponseMsg))
}

func BuildResponseOKWithContent(p_Context *gin.Context, p_ResponseCode string, p_ResponseMsg string, p_Payload interface{}) {
	var result dto.ResponseDataContent
	result.ResponseData = BuildResponseData(p_ResponseCode, p_ResponseMsg)
	result.Content = p_Payload
	p_Context.JSON(http.StatusOK, result)
}

func BuildResponseInternalServerError(p_Context *gin.Context, p_ResponseCode string, p_ResponseMsg string) {
	p_Context.JSON(http.StatusInternalServerError, BuildResponseData(p_ResponseCode, p_ResponseMsg))
}

func BuildResponseBadRequest(p_Context *gin.Context, p_ResponseCode string, p_ResponseMsg string) {
	p_Context.JSON(http.StatusBadRequest, BuildResponseData(p_ResponseCode, p_ResponseMsg))
}

func BuildResponseConflict(p_Context *gin.Context, p_ResponseCode string, p_ResponseMsg string) {
	p_Context.JSON(http.StatusConflict, BuildResponseData(p_ResponseCode, p_ResponseMsg))
}

func BuildResponseData(p_ResponseCode string, p_ResponseMsg string) dto.ResponseData {
	var result dto.ResponseData
	result.ResponseCode = p_ResponseCode
	result.ResponseMsg = p_ResponseMsg
	return result
}
