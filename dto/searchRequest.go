package dto

type SearchRequest struct {
	SearchCriteria []SearchCriteria
	SortBy         string
	Page           int
	Limit          int
}
