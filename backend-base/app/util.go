package app

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"regexp"
	"restaurant/backend-base/entity"
	"restaurant/backend-base/logger"
)

var IsValidUsername = regexp.MustCompile("([a-zA-Z0-9](_|-| )[a-zA-Z0-9])*").MatchString

func ConvertToJson(m proto.Message) string {
	data, err := entity.Marshaler.MarshalToString(m)
	if err != nil {
		logger.Logger.Error(err)
		return "{}"
	}
	return data
}
func ConvertBaseRequestToJson(m entity.BaseRequest) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		logger.Logger.Error(err)
		return nil
	}
	return data
}

func ConvertInterfaceToString(data interface{}) string {
	if data == nil {
		return ""
	}
	return data.(string)
}

func IsLetter(s string) bool {
	for _, r := range s {
		if r != ' ' && (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
