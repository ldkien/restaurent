package entity

import (
	"encoding/json"
	pb "restaurant/backend-entity/entities"
)

type BaseRequest struct {
	Common *pb.Common
	Data   *json.RawMessage
}
