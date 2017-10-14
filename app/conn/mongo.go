package conn

import (
	"github.com/go-bongo/bongo"
	"github.com/spf13/viper"
)

var conn *bongo.Connection
var err error

func GetConnection() *bongo.Connection {
	if conn == nil {
		conn, err = bongo.Connect(&bongo.Config{
			ConnectionString: viper.GetString("databases.mongodb.uri"),
			Database:         viper.GetString("databases.mongodb.dbname"),
		})
		if err != nil {
			panic(err)
		}
		return conn
	}
	return conn
}
