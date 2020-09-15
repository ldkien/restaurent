package table

import (
	"github.com/gocql/gocql"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	pb "restaurant/backend-entity/entities"
)

func GetTableDetail(tableId string) *pb.Table {
	query := "SELECT * FROM table_data where table_id = ?"
	iterator := cassandra.Session.
		Query(query, tableId).Consistency(gocql.One).Iter()
	m := map[string]interface{}{}

	for iterator.MapScan(m) {
		return &pb.Table{
			TableId:   app.ConvertInterfaceToString(m["table_id"]),
			TableName: app.ConvertInterfaceToString(m["table_name"]),
			TableDesc: app.ConvertInterfaceToString(m["description"]),
		}
	}
	return nil
}
