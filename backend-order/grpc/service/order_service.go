package service

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	"restaurant/backend-base/database/dish"
	"restaurant/backend-base/database/table"
	backendError "restaurant/backend-base/error"
	"restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"time"
)

func Order(in *pb.OrderRequest) (*pb.BaseResponse, error) {
	error, response := validateData(in)

	if error.ErrorCode != backendError.SUCCESS {
		return &pb.BaseResponse{
			Error: error,
			Data:  app.ConvertDataAny(in),
		}, nil
	}

	id := createOrderRequest(in)
	if id == nil {
		error = backendError.GetError(backendError.DB_ERROR)
	}

	response.Id = *id
	return &pb.BaseResponse{
		Error: error,
		Data:  app.ConvertDataAny(response),
	}, nil
}

func validateData(input *pb.OrderRequest) (*pb.Error, *pb.OrderResponse) {

	if input == nil || len(input.Dishes) <= 0 ||
		len(input.Table.TableId) <= 0 || len(input.Note) > 1000 || !validateQuantity(input.Dishes) {
		return backendError.GetError(backendError.INVALID_PARAMS), nil
	}

	table := table.GetTableDetail(input.Table.TableId)

	if table == nil {
		logger.Logger.Error("No exist table: " + input.Table.TableId)
		return backendError.GetError(backendError.DB_NO_RECORD), nil
	}

	dishes := dish.GetDishByIds(input.Dishes)

	if len(dishes) != len(input.Dishes) {
		logger.Logger.Error("Some one dishes exist")
		return backendError.GetError(backendError.DB_NO_RECORD), nil
	}

	response := &pb.OrderResponse{
		Table:  table,
		Dishes: dishes,
	}
	return backendError.GetError(backendError.SUCCESS), response
}

func validateQuantity(dishes []*pb.Dish) bool {
	for _, item := range dishes {
		if item.Quantity <= 0 {
			return false
		}
	}
	return true
}

func createOrderRequest(input *pb.OrderRequest) *string {
	currentTime := time.Now()
	var sql = "INSERT INTO order_month_%d(id,table_id,username,data," +
		"create_date,status,comment,update_by) VALUES(?,?,?,?,?,?,?,?)"
	sql = fmt.Sprintf(sql, int(currentTime.Month()))
	id := uuid.New().String()
	if err := cassandra.Session.Query(sql, id,
		input.Table.TableId, input.Common.User.Username,
		app.ConvertObjectToStringJson(input.Dishes), currentTime, pb.STATUS_DELIVERY_RECEIVED.Number(),
		input.Note, input.Common.User.Username).Exec(); err != nil {
		logger.Logger.Error("Insert order error", err)
		return nil
	}

	return &id
}

func UpdateOrder(in *pb.OrderRequest) (*pb.BaseResponse, error) {

}

func validateUpdateOrder(input *pb.OrderRequest) (*pb.Error, *pb.OrderResponse) {
	if input == nil || len(input.Dishes) <= 0 ||
		len(input.Id) <= 0 || len(input.Note) > 1000 || !validateQuantity(input.Dishes) {
		return backendError.GetError(backendError.INVALID_PARAMS), nil
	}

	dishes := dish.GetDishByIds(input.Dishes)

	if len(dishes) != len(input.Dishes) {
		logger.Logger.Error("Some one dishes exist")
		return backendError.GetError(backendError.DB_NO_RECORD), nil
	}

	input.Dishes = dishes

	return backendError.GetError(backendError.SUCCESS), getOrderById(input)

}

func getOrderById(input *pb.OrderRequest) *pb.OrderResponse {
	currentTime := time.Now()
	var sql = "SELECT * FROM order_month_%d WHERE id = ?"
	sql = fmt.Sprintf(sql, int(currentTime.Month()))
	iterator := cassandra.Session.
		Query(sql, input.Id).Consistency(gocql.One).Iter()
	m := map[string]interface{}{}

	for iterator.MapScan(m) {
		input.Table.TableId = app.ConvertInterfaceToString(m["table_id"])
		return &pb.OrderResponse{
			Id:           input.Id,
			Table:        input.Table,
			Dishes:       input.Dishes,
			Note:         input.Note,
			Status:       input.Status,
			TimeDelivery: 0,
		}
	}

	return nil
}
