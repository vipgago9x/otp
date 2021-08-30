package models

type ResponseModel struct {
	Error ErrorModel `json:"error"`
	Data  DataModel  `json:"data"`
}
