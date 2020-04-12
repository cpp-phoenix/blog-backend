package dto

type SearchCriteria struct {
	MustNotValueFilters []MustNotValueFilters
	MustValueFilters    []MustValueFilters
}
