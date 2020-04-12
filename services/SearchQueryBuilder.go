package services

import "blog_backend/dto"

func searchRequestBuilderForUserName(userName string) dto.SearchRequest {
	var stringValueFilters dto.StringValueFilters
	stringValueFilters.FieldName = "username"
	stringValueFilters.Values = append(stringValueFilters.Values, userName)

	var mustValueFilters dto.MustValueFilters
	mustValueFilters.StringValueFilters = append(mustValueFilters.StringValueFilters, stringValueFilters)

	var searchCriteria dto.SearchCriteria
	searchCriteria.MustValueFilters = append(searchCriteria.MustValueFilters, mustValueFilters)

	var searchRequest dto.SearchRequest
	searchRequest.SearchCriteria = append(searchRequest.SearchCriteria, searchCriteria)
	return searchRequest
}

func searchRequestBuilderForEmailAddress(email string) dto.SearchRequest {
	var stringValueFilters dto.StringValueFilters
	stringValueFilters.FieldName = "email"
	stringValueFilters.Values = append(stringValueFilters.Values, email)

	var mustValueFilters dto.MustValueFilters
	mustValueFilters.StringValueFilters = append(mustValueFilters.StringValueFilters, stringValueFilters)

	var searchCriteria dto.SearchCriteria
	searchCriteria.MustValueFilters = append(searchCriteria.MustValueFilters, mustValueFilters)

	var searchRequest dto.SearchRequest
	searchRequest.SearchCriteria = append(searchRequest.SearchCriteria, searchCriteria)
	return searchRequest
}
