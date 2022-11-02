package model

type PaginationResponse struct {
	Page       int `json:"page"`
	TotalPages int `json:"totalPages"`
}
