package util

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Content         interface{} `json:"content"`
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMsg"`
}

func ResponseJSON(c *gin.Context, status int, payload interface{}) {
	var result ResponseData
	result.ResponseCode = status
	result.ResponseMessage = "Success"
	result.Content = payload

	c.JSON(200, result)
}
