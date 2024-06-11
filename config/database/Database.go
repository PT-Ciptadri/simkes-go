package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	const MYSQL_DSN = "root:@tcp(127.0.0.1:3306)/simkes?charset=utf8mb4&parseTime=True&loc=Local"

	DSN := MYSQL_DSN

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")

}
