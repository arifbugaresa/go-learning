package common

type ListRequest struct {
	ID int64 `json:"id"`
	Filter
	Sort
}

type Filter struct {
	Page  *int64 `json:"page"`
	Limit *int64 `json:"limit"`
}

type Sort struct {
	SortField *string `json:"field"`
	SortOrder *string `json:"order"`
}
