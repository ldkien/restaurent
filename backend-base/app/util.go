package app

import (
	"github.com/golang/protobuf/proto"
	"restaurant/backend-base/entity"
	"restaurant/backend-base/logger"
)

func ConvertToJson(m proto.Message) string {
	data, err := entity.Marshaler.MarshalToString(m)
	if err != nil {
		logger.Logger.Error(err)
		return "{}"
	}
	return data
}

func ConvertInterfaceToString(data interface{}) string {
	if data == nil {
		return ""
	}
	return data.(string)
}
