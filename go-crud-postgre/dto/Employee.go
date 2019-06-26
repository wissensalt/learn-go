package dto

type Employee struct {
	Id      int     `json:"id"`
	Code    string  `json:"code"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salary  float64 `json:salary`
}
