package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func getStringFilterEqObject(stringValueFilters dto.StringValueFilters) dto.BsonQueryDTO {
	var object dto.BsonQueryDTO
	object.Key = stringValueFilters.FieldName
	var nestedObject dto.BsonQueryDTO
	if len(stringValueFilters.Values) > 1 {
		nestedObject.Key = properties.IN_FILTER
		nestedObject.Value = stringValueFilters.Values
	} else {
		nestedObject.Key = properties.EQ_FILTER
		nestedObject.Value = stringValueFilters.Values[0]
	}
	object.Value = nestedObject
	return object
}

func getStringFilterNotEqObject(stringValueFilters dto.StringValueFilters) dto.BsonQueryDTO {
	var object dto.BsonQueryDTO
	object.Key = stringValueFilters.FieldName
	var nestedObject dto.BsonQueryDTO
	if len(stringValueFilters.Values) > 1 {
		nestedObject.Key = properties.NOT_IN_FILTER
		nestedObject.Value = stringValueFilters.Values
	} else {
		nestedObject.Key = properties.NOT_EQUAL_FILTER
		nestedObject.Value = stringValueFilters.Values[0]
	}
	object.Value = nestedObject
	return object
}

func getLongFilterEqObject(intValueFilters dto.IntValueFilters) dto.BsonQueryDTO {
	var object dto.BsonQueryDTO
	object.Key = intValueFilters.FieldName
	var nestedObject dto.BsonQueryDTO
	if len(intValueFilters.Values) > 1 {
		nestedObject.Key = properties.IN_FILTER
		nestedObject.Value = intValueFilters.Values
	} else {
		nestedObject.Key = properties.EQ_FILTER
		nestedObject.Value = intValueFilters.Values[0]
	}
	object.Value = nestedObject
	return object
}

func getLongFilterNotEqObject(intValueFilters dto.IntValueFilters) dto.BsonQueryDTO {
	var object dto.BsonQueryDTO
	object.Key = intValueFilters.FieldName
	var nestedObject dto.BsonQueryDTO
	if len(intValueFilters.Values) > 1 {
		nestedObject.Key = properties.NOT_IN_FILTER
		nestedObject.Value = intValueFilters.Values
	} else {
		nestedObject.Key = properties.NOT_EQUAL_FILTER
		nestedObject.Value = intValueFilters.Values[0]
	}
	object.Value = nestedObject
	return object
}

func longValueFilterIteratorNotEqObject(intValueFilters []dto.IntValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(intValueFilters) > 1 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(intValueFilters); i++ {
			bsonArray = append(bsonArray, getLongFilterNotEqObject(intValueFilters[i]))
		}
		filterObject.Value = bsonArray
	} else {
		filterObject = getLongFilterNotEqObject(intValueFilters[0])
	}
	return filterObject
}

func stringValueFilterIteratorNotEqObject(stringValueFilters []dto.StringValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(stringValueFilters) > 1 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(stringValueFilters); i++ {
			bsonArray = append(bsonArray, getStringFilterNotEqObject(stringValueFilters[i]))
		}
		filterObject.Value = bsonArray
	} else {
		filterObject = getStringFilterNotEqObject(stringValueFilters[0])
	}
	return filterObject
}

func longValueFilterIteratorForEqObject(intValueFilters []dto.IntValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(intValueFilters) > 1 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(intValueFilters); i++ {
			bsonArray = append(bsonArray, getLongFilterEqObject(intValueFilters[i]))
		}
		filterObject.Value = bsonArray
	} else {
		filterObject = getLongFilterEqObject(intValueFilters[0])
	}
	return filterObject
}

func stringValueFilterIteratorForEqObject(stringValueFilters []dto.StringValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(stringValueFilters) > 1 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(stringValueFilters); i++ {
			bsonArray = append(bsonArray, getStringFilterEqObject(stringValueFilters[i]))
		}
		filterObject.Value = bsonArray
	} else {
		filterObject = getStringFilterEqObject(stringValueFilters[0])
	}
	return filterObject
}

func returnMustFilterObject(mustValueFilterObject dto.MustValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(mustValueFilterObject.IntValueFilters) > 0 && len(mustValueFilterObject.StringValueFilters) > 0 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		bsonArray = append(bsonArray, longValueFilterIteratorForEqObject(mustValueFilterObject.IntValueFilters))
		bsonArray = append(bsonArray, stringValueFilterIteratorForEqObject(mustValueFilterObject.StringValueFilters))
		filterObject.Value = bsonArray
	} else if len(mustValueFilterObject.IntValueFilters) > 0 {
		filterObject = longValueFilterIteratorForEqObject(mustValueFilterObject.IntValueFilters)
	} else if len(mustValueFilterObject.StringValueFilters) > 0 {
		filterObject = stringValueFilterIteratorForEqObject(mustValueFilterObject.StringValueFilters)
	}
	return filterObject
}

func returnMustNotFilterObject(mustNotValueFilterObject dto.MustNotValueFilters) dto.BsonQueryDTO {
	var filterObject dto.BsonQueryDTO
	if len(mustNotValueFilterObject.IntValueFilters) > 0 && len(mustNotValueFilterObject.StringValueFilters) > 0 {
		filterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		bsonArray = append(bsonArray, longValueFilterIteratorNotEqObject(mustNotValueFilterObject.IntValueFilters))
		bsonArray = append(bsonArray, stringValueFilterIteratorNotEqObject(mustNotValueFilterObject.StringValueFilters))
		filterObject.Value = bsonArray
	} else if len(mustNotValueFilterObject.IntValueFilters) > 0 {
		filterObject = longValueFilterIteratorNotEqObject(mustNotValueFilterObject.IntValueFilters)
	} else if len(mustNotValueFilterObject.StringValueFilters) > 0 {
		filterObject = stringValueFilterIteratorNotEqObject(mustNotValueFilterObject.StringValueFilters)
	}
	return filterObject
}

func searchCriteriaToMustQueryBinder(mustValueFilters []dto.MustValueFilters) dto.BsonQueryDTO {
	var mustFilterObject dto.BsonQueryDTO
	if len(mustValueFilters) > 1 {
		mustFilterObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(mustValueFilters); i++ {
			bsonArray = append(bsonArray, returnMustFilterObject(mustValueFilters[i]))
		}
		mustFilterObject.Value = bsonArray
	} else {
		mustFilterObject = returnMustFilterObject(mustValueFilters[0])
	}
	return mustFilterObject
}

func searchCriteriaToMustNotQueryBinder(mustNotValueFilters []dto.MustNotValueFilters) dto.BsonQueryDTO {
	var mustNotFilters dto.BsonQueryDTO
	if len(mustNotValueFilters) > 1 {
		mustNotFilters.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(mustNotValueFilters); i++ {
			bsonArray = append(bsonArray, returnMustNotFilterObject(mustNotValueFilters[i]))
		}
		mustNotFilters.Value = bsonArray
	} else {
		mustNotFilters = returnMustNotFilterObject(mustNotValueFilters[0])
	}
	return mustNotFilters
}

func returnSeachCriteriaObject(searchCriteria dto.SearchCriteria) dto.BsonQueryDTO {
	var searchCriteriaBsonObject dto.BsonQueryDTO
	if len(searchCriteria.MustNotValueFilters) > 0 && len(searchCriteria.MustValueFilters) > 0 {
		searchCriteriaBsonObject.Key = properties.AND_FILTER
		var bsonArray []dto.BsonQueryDTO
		bsonArray = append(bsonArray, searchCriteriaToMustQueryBinder(searchCriteria.MustValueFilters))
		bsonArray = append(bsonArray, searchCriteriaToMustNotQueryBinder(searchCriteria.MustNotValueFilters))
		searchCriteriaBsonObject.Value = bsonArray
	} else if len(searchCriteria.MustValueFilters) > 0 {
		searchCriteriaBsonObject = searchCriteriaToMustQueryBinder(searchCriteria.MustValueFilters)
	} else if len(searchCriteria.MustNotValueFilters) > 0 {
		searchCriteriaBsonObject = searchCriteriaToMustNotQueryBinder(searchCriteria.MustNotValueFilters)
	}
	return searchCriteriaBsonObject
}

func createDTOToBsonRequest(searchRequest dto.SearchRequest) bson.M {
	var query dto.BsonQueryDTO
	if len(searchRequest.SearchCriteria) == 0 {
		//Do something!!
		//Error handeling should be applied here!!
	} else if len(searchRequest.SearchCriteria) > 1 {
		query.Key = properties.OR_FILTER
		var bsonArray []dto.BsonQueryDTO
		for i := 0; i < len(searchRequest.SearchCriteria); i++ {
			bsonArray = append(bsonArray, returnSeachCriteriaObject(searchRequest.SearchCriteria[i]))
		}
		query.Value = bsonArray
	} else {
		query = returnSeachCriteriaObject(searchRequest.SearchCriteria[0])
	}
	var bsonMap bson.M
	fmt.Println(query)
	foo_marshalled, err := json.Marshal(query)
	err = json.Unmarshal([]byte(foo_marshalled), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}
