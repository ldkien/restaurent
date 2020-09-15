package dish

import (
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	pb "restaurant/backend-entity/entities"
)

func GetDishByIds(dishes []*pb.Dish) []*pb.Dish {
	var result []*pb.Dish
	query := "SELECT * FROM dish_data WHERE dish_id IN ("

	for index, item := range dishes {
		if index == len(dishes)-1 {
			query += "'" + item.DishId + "')"
		} else {
			query += "'" + item.DishId + "',"
		}
	}
	iterator := cassandra.Session.
		Query(query).Iter()
	m := map[string]interface{}{}

	for iterator.MapScan(m) {
		var dish = &pb.Dish{
			DishId:   app.ConvertInterfaceToString(m["dish_id"]),
			DishName: app.ConvertInterfaceToString(m["dish_name"]),
			DishNote: app.ConvertInterfaceToString(m["description"]),
		}
		result = append(result, dish)
		m = map[string]interface{}{}
	}
	return result
}
