package services

import (
	"blog_backend/dto"
)

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

func searchRequestBuilderForPostId(postId int) dto.SearchRequest {
	var longValueFilters dto.IntValueFilters
	longValueFilters.FieldName = "postid"
	longValueFilters.Values = append(longValueFilters.Values, (postId))

	var mustValueFilters dto.MustValueFilters
	mustValueFilters.IntValueFilters = append(mustValueFilters.IntValueFilters, longValueFilters)

	var searchCriteria dto.SearchCriteria
	searchCriteria.MustValueFilters = append(searchCriteria.MustValueFilters, mustValueFilters)

	var searchRequest dto.SearchRequest
	searchRequest.SearchCriteria = append(searchRequest.SearchCriteria, searchCriteria)
	return searchRequest
}
