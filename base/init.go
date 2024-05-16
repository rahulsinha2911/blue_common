package base

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	MysqlClient *gorm.DB
)

func Init() {
	godotenv.Load()

	// Init Users DB Connection
	mysqlClient, err := GetMySqlDBConfig()
	if err != nil {
		panic(err)
	}
	MysqlClient = mysqlClient
}
