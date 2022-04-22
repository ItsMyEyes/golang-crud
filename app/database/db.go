package database

import (
	"crud_v2/app/enviroment"
	"crud_v2/entity"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	initDB()
}

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var Connector *gorm.DB

//Connect creates MySQL connection
func connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

var GetConnectionString = func(config *Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}

func initDB() {
	config := Config{
		ServerName: enviroment.Get("DB_HOST"),
		User:       enviroment.Get("DB_USER"),
		Password:   enviroment.Get("DB_PASSWORD"),
		DB:         enviroment.Get("DB_NAME"),
	}
	connectionString := GetConnectionString(&config)
	err := connect(connectionString)
	if err != nil {
		log.Println("Connection failed!!")
		log.Fatal(err)
	}

	Connector.LogMode(true)
	Connector.AutoMigrate(&entity.Todo{}, &entity.User{})
}
