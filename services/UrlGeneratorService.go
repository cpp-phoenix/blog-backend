package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func createBsonObjectForDataSaving(user dto.UserDetails) bson.M {
	query := make(map[string]interface{})

	internalquery := make(map[string]interface{})
	internalquery["otp"] = user.Otp
	internalquery["resetrequesttimstamp"] = user.ResetRequestTimstamp

	query["$set"] = internalquery
	var bsonMap bson.M
	foo_marshalled, _ := json.Marshal(query)
	err := json.Unmarshal([]byte(string(foo_marshalled)), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}

func GenerateRandomNumber(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GenerateUrl(user dto.UserDetails) string {
	if user.Otp != 0 && ((time.Now().Unix()-user.ResetRequestTimstamp)/3600) < (properties.HOURS_TO_OTP_EXPIRE) {
		return "userName=" + user.UserName + "&otp=" + strconv.Itoa(user.Otp)
	}
	user.Otp = GenerateRandomNumber(1000, 9999)
	user.ResetRequestTimstamp = time.Now().Unix()

	searchRequest := searchRequestBuilderForUserName(user.UserName)
	//Updating new otp to DB
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest, createBsonObjectForDataSaving(user))
	if status == 3012 {
		return "3012"
	}
	return "userName=" + user.UserName + "&otp=" + strconv.Itoa(user.Otp)
}
