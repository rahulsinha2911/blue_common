package base

import (
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMySqlDBConfig() (*gorm.DB, error) {
	idle_connection, err := strconv.Atoi(os.Getenv("BLUE_FERRY_DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		panic(err)
	}
	max_connection, err := strconv.Atoi(os.Getenv("BLUE_FERRY_DB_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("BLUE_FERRY_DB_HOST")
	dbPort := os.Getenv("BLUE_FERRY_DB_PORT")
	dbUser := os.Getenv("BLUE_FERRY_DB_USERNAME")
	dbPassword := os.Getenv("BLUE_FERRY_DB_PASSWORD")
	dbName := os.Getenv("BLUE_FERRY_DB_SCHEMA")

	// Now you can establish the database connection using GORM
	// For example, for MySQL:
	dbParams := "?parseTime=true&charset=utf8mb4&loc=Asia%2FKolkata"
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + dbParams

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	sqlDb, err := db.DB()
	sqlDb.SetMaxIdleConns(idle_connection)
	sqlDb.SetMaxOpenConns(max_connection)
	sqlDb.SetConnMaxLifetime(3 * time.Minute)

	return db, sqlDb.Ping()
}
