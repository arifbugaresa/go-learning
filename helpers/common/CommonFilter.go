package common

type SearchAndFilter struct {
	ID       int64       `json:"id"`
	SearchBy interface{} `json:"search"`
	Filter
}

type Filter struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
