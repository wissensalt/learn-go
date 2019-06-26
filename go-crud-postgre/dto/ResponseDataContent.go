package dto

type ResponseDataContent struct {
	ResponseData ResponseData `json:"responseCode"`
	Content      interface{}  `json:content`
}
