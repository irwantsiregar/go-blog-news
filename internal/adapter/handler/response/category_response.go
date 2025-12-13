package response

type SuccessCategoryResponse struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Slug string `json:"slug"`
	CreatedByName string `json:"created_by_name"`
}

type DefaultSuccessResponse struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination,omitempty"`
}

type PaginationResponse struct {
	TotalRecords	int `json:"total_records"`
	Page			int `json:"page"`
	PerPage			int `json:"per_page"`
	TotalPages		int `json:"total_pages"`
}