package common

type SearchAndFilter struct {
	ID     int64       `form:"id"`
	Search interface{} `form:"search"`
	Filter
}

type Filter struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
