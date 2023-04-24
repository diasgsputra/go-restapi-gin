package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var user = os.Getenv("MYSQL_USER")
var password = os.Getenv("MYSQL_PASSWORD")
var host = os.Getenv("MYSQL_HOST")
var port = os.Getenv("MYSQL_PORT")
var database = os.Getenv("MYSQL_DBNAME")
// var user = "root"
// var password = ""
// var host = "localhost"
// var port = "3306"
// var database = "todo4"


func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(
		user+":"+password+"@tcp("+host+":"+port+")/"+database))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Activities{})
	database.AutoMigrate(&Todo{})

	DB = database
}
