package model

type StandardResponse struct {
	Status StatusResponse `json:"status"`
	Result interface{}    `json:"result"`
	Page   interface{}    `json:"page"`
}
