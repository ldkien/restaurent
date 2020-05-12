package service

import (
	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	backendError "restaurant/backend-base/error"
	"restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
)

func Order(in *pb.OrderRequest) (*pb.BaseResponse, error) {
	response := validateData(in)

	if response.Error.ErrorCode != backendError.SUCCESS {
		return response, nil
	}

	return response, nil
}

func validateData(input *pb.OrderRequest) *pb.BaseResponse {

	if input == nil || len(input.DishId) <= 0 ||
		len(input.TableId) <= 0 || len(input.Note) > 1000 || input.Num <= 0 {
		return &pb.BaseResponse{
			Error: backendError.GetError(backendError.INVALID_PARAMS),
			Data:  nil,
		}
	}

	table := getTableDetail(input.TableId)

	if table == nil {
		logger.Logger.Error("No exist table: " + input.TableId)
		return &pb.BaseResponse{
			Error: backendError.GetError(backendError.DB_NO_RECORD),
			Data:  nil,
		}
	}

	response := &pb.OrderResponse{
		TableId:   table.TableId,
		TableName: table.TableName,
		TableDesc: table.Desc,
	}
	data, err := ptypes.MarshalAny(response)
	if err != nil {
		logger.Logger.Error("Cannot parse any", err)
		return &pb.BaseResponse{
			Error: backendError.GetError(backendError.SYSTEM_ERROR),
			Data:  nil,
		}
	}
	return &pb.BaseResponse{
		Error: backendError.GetError(backendError.SUCCESS),
		Data:  data,
	}

}

func getTableDetail(tableId string) *pb.Table {
	query := "SELECT * FROM table_data where tableId = ?"
	iterator := cassandra.Session.
		Query(query, tableId).Consistency(gocql.One).Iter()
	m := map[string]interface{}{}

	for iterator.MapScan(m) {
		return &pb.Table{
			TableId:   app.ConvertInterfaceToString(m["tableid"]),
			TableName: app.ConvertInterfaceToString(m["tablename"]),
			Desc:      app.ConvertInterfaceToString(m["desc"]),
		}
	}
	return nil
}
