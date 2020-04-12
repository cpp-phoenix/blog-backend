package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func getStringFilterEqObject(stringValueFilters dto.StringValueFilters) map[string]interface{} {
	object := make(map[string]interface{})
	nestedObject := make(map[string]interface{})
	if len(stringValueFilters.Values) > 1 {
		//nestedObject.Key = properties.IN_FILTER
		nestedObject[properties.IN_FILTER] = stringValueFilters.Values
	} else {
		//nestedObject.Key = properties.EQ_FILTER
		nestedObject[properties.EQ_FILTER] = stringValueFilters.Values[0]
	}
	object[stringValueFilters.FieldName] = nestedObject
	return object
}

func getStringFilterNotEqObject(stringValueFilters dto.StringValueFilters) map[string]interface{} {
	object := make(map[string]interface{})
	nestedObject := make(map[string]interface{})
	if len(stringValueFilters.Values) > 1 {
		//nestedObject.Key = properties.NOT_IN_FILTER
		nestedObject[properties.NOT_IN_FILTER] = stringValueFilters.Values
	} else {
		//nestedObject.Key = properties.NOT_EQUAL_FILTER
		nestedObject[properties.NOT_EQUAL_FILTER] = stringValueFilters.Values[0]
	}
	object[stringValueFilters.FieldName] = nestedObject
	return object
}

func getLongFilterEqObject(intValueFilters dto.IntValueFilters) map[string]interface{} {
	object := make(map[string]interface{})
	nestedObject := make(map[string]interface{})
	if len(intValueFilters.Values) > 1 {
		//nestedObject.Key = properties.IN_FILTER
		nestedObject[properties.IN_FILTER] = intValueFilters.Values
	} else {
		//nestedObject.Key = properties.EQ_FILTER
		nestedObject[properties.EQ_FILTER] = intValueFilters.Values[0]
	}
	object[intValueFilters.FieldName] = nestedObject
	return object
}

func getLongFilterNotEqObject(intValueFilters dto.IntValueFilters) map[string]interface{} {
	object := make(map[string]interface{})
	nestedObject := make(map[string]interface{})
	if len(intValueFilters.Values) > 1 {
		//nestedObject.Key = properties.NOT_IN_FILTER
		nestedObject[properties.NOT_IN_FILTER] = intValueFilters.Values
	} else {
		//nestedObject.Key = properties.NOT_EQUAL_FILTER
		nestedObject[properties.NOT_EQUAL_FILTER] = intValueFilters.Values[0]
	}
	object[intValueFilters.FieldName] = nestedObject
	return object
}

func longValueFilterIteratorNotEqObject(intValueFilters []dto.IntValueFilters) map[string]interface{} {
	filterObject := make(map[string]interface{})
	if len(intValueFilters) > 1 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(intValueFilters); i++ {
			bsonArray = append(bsonArray, getLongFilterNotEqObject(intValueFilters[i]))
		}
		filterObject[properties.AND_FILTER] = bsonArray
	} else {
		filterObject = getLongFilterNotEqObject(intValueFilters[0])
	}
	return filterObject
}

func stringValueFilterIteratorNotEqObject(stringValueFilters []dto.StringValueFilters) map[string]interface{} {
	filterObject := make(map[string]interface{})
	if len(stringValueFilters) > 1 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(stringValueFilters); i++ {
			bsonArray = append(bsonArray, getStringFilterNotEqObject(stringValueFilters[i]))
		}
		filterObject[properties.AND_FILTER] = bsonArray
	} else {
		filterObject = getStringFilterNotEqObject(stringValueFilters[0])
	}
	return filterObject
}

func longValueFilterIteratorForEqObject(intValueFilters []dto.IntValueFilters) map[string]interface{} {
	var filterObject map[string]interface{}
	if len(intValueFilters) > 1 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(intValueFilters); i++ {
			bsonArray = append(bsonArray, getLongFilterEqObject(intValueFilters[i]))
		}
		filterObject := make(map[string]interface{})
		filterObject[properties.AND_FILTER] = bsonArray
	} else {
		filterObject = getLongFilterEqObject(intValueFilters[0])
	}
	return filterObject
}

func stringValueFilterIteratorForEqObject(stringValueFilters []dto.StringValueFilters) map[string]interface{} {
	filterObject := make(map[string]interface{})
	if len(stringValueFilters) > 1 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(stringValueFilters); i++ {
			bsonArray = append(bsonArray, getStringFilterEqObject(stringValueFilters[i]))
		}

		filterObject[properties.AND_FILTER] = bsonArray
	} else {
		filterObject = getStringFilterEqObject(stringValueFilters[0])
	}
	return filterObject
}

func returnMustFilterObject(mustValueFilterObject dto.MustValueFilters) map[string]interface{} {
	filterObject := make(map[string]interface{})
	if len(mustValueFilterObject.IntValueFilters) > 0 && len(mustValueFilterObject.StringValueFilters) > 0 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		bsonArray = append(bsonArray, longValueFilterIteratorForEqObject(mustValueFilterObject.IntValueFilters))
		bsonArray = append(bsonArray, stringValueFilterIteratorForEqObject(mustValueFilterObject.StringValueFilters))

		filterObject[properties.AND_FILTER] = bsonArray
	} else if len(mustValueFilterObject.IntValueFilters) > 0 {
		filterObject = longValueFilterIteratorForEqObject(mustValueFilterObject.IntValueFilters)
	} else if len(mustValueFilterObject.StringValueFilters) > 0 {
		filterObject = stringValueFilterIteratorForEqObject(mustValueFilterObject.StringValueFilters)
	}
	return filterObject
}

func returnMustNotFilterObject(mustNotValueFilterObject dto.MustNotValueFilters) map[string]interface{} {
	filterObject := make(map[string]interface{})
	if len(mustNotValueFilterObject.IntValueFilters) > 0 && len(mustNotValueFilterObject.StringValueFilters) > 0 {
		//filterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		bsonArray = append(bsonArray, longValueFilterIteratorNotEqObject(mustNotValueFilterObject.IntValueFilters))
		bsonArray = append(bsonArray, stringValueFilterIteratorNotEqObject(mustNotValueFilterObject.StringValueFilters))

		filterObject[properties.AND_FILTER] = bsonArray
	} else if len(mustNotValueFilterObject.IntValueFilters) > 0 {
		filterObject = longValueFilterIteratorNotEqObject(mustNotValueFilterObject.IntValueFilters)
	} else if len(mustNotValueFilterObject.StringValueFilters) > 0 {
		filterObject = stringValueFilterIteratorNotEqObject(mustNotValueFilterObject.StringValueFilters)
	}
	return filterObject
}

func searchCriteriaToMustQueryBinder(mustValueFilters []dto.MustValueFilters) map[string]interface{} {
	mustFilterObject := make(map[string]interface{})
	if len(mustValueFilters) > 1 {
		//mustFilterObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(mustValueFilters); i++ {
			bsonArray = append(bsonArray, returnMustFilterObject(mustValueFilters[i]))
		}

		mustFilterObject[properties.AND_FILTER] = bsonArray
	} else {
		mustFilterObject = returnMustFilterObject(mustValueFilters[0])
	}
	return mustFilterObject
}

func searchCriteriaToMustNotQueryBinder(mustNotValueFilters []dto.MustNotValueFilters) map[string]interface{} {
	mustNotFilters := make(map[string]interface{})
	if len(mustNotValueFilters) > 1 {
		//mustNotFilters.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(mustNotValueFilters); i++ {
			bsonArray = append(bsonArray, returnMustNotFilterObject(mustNotValueFilters[i]))
		}

		mustNotFilters[properties.AND_FILTER] = bsonArray
	} else {
		mustNotFilters = returnMustNotFilterObject(mustNotValueFilters[0])
	}
	return mustNotFilters
}

func returnSeachCriteriaObject(searchCriteria dto.SearchCriteria) map[string]interface{} {
	searchCriteriaBsonObject := make(map[string]interface{})
	if len(searchCriteria.MustNotValueFilters) > 0 && len(searchCriteria.MustValueFilters) > 0 {
		//searchCriteriaBsonObject.Key = properties.AND_FILTER
		var bsonArray []map[string]interface{}
		bsonArray = append(bsonArray, searchCriteriaToMustQueryBinder(searchCriteria.MustValueFilters))
		bsonArray = append(bsonArray, searchCriteriaToMustNotQueryBinder(searchCriteria.MustNotValueFilters))

		searchCriteriaBsonObject[properties.AND_FILTER] = bsonArray
	} else if len(searchCriteria.MustValueFilters) > 0 {
		searchCriteriaBsonObject = searchCriteriaToMustQueryBinder(searchCriteria.MustValueFilters)
	} else if len(searchCriteria.MustNotValueFilters) > 0 {
		searchCriteriaBsonObject = searchCriteriaToMustNotQueryBinder(searchCriteria.MustNotValueFilters)
	}
	return searchCriteriaBsonObject
}

func createDTOToBsonRequest(searchRequest dto.SearchRequest) bson.M {
	query := make(map[string]interface{})
	if len(searchRequest.SearchCriteria) == 0 {
		//Do something!!
		//Error handeling should be applied here!!
	} else if len(searchRequest.SearchCriteria) > 1 {
		//query.Key = properties.OR_FILTER
		var bsonArray []map[string]interface{}
		for i := 0; i < len(searchRequest.SearchCriteria); i++ {
			bsonArray = append(bsonArray, returnSeachCriteriaObject(searchRequest.SearchCriteria[i]))
		}
		query[properties.OR_FILTER] = bsonArray
	} else {
		query = returnSeachCriteriaObject(searchRequest.SearchCriteria[0])
	}
	var bsonMap bson.M
	foo_marshalled, _ := json.Marshal(query)
	err := json.Unmarshal([]byte(string(foo_marshalled)), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}
