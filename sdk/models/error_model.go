package models

type ErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
