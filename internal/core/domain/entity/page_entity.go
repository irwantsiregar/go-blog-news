package entity

type Page struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	First      int `json:"first"`
	Last       int `json:"last"`
}