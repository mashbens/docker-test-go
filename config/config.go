package config

import (
	"fmt"
	"postgresql-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "arif1412"
	DB_NAME     = "test_db"
	Db_PORT     = "5432"
)

var DB *gorm.DB

var err error

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, Db_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}

func DBManager() *gorm.DB {
	return DB
}
