package cassandra

import (
	"github.com/gocql/gocql"
	"restaurant/backend-base/app"
	log "restaurant/backend-base/logger"
)
var Session *gocql.Session

func init() {

	var err error
	cluster := gocql.NewCluster(app.GlobalConfig.Cassandra.Host)
	cluster.Keyspace = app.GlobalConfig.Cassandra.Keyspace
	cluster.Port = app.GlobalConfig.Cassandra.Port
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}
	log.Logger.Info("cassandra init done")
}

func Close()  {
	log.Logger.Info("Close session cassandra")
	Session.Close()
}