package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	envHostName := "DB_HOST"
	envUserName := "DB_USER"
	envPassword := "DB_PASSWORD"
	envDBName := "DB_NAME"

	host := os.Getenv(envHostName)
	userName := os.Getenv(envUserName)
	password := os.Getenv(envPassword)
	dbName := os.Getenv(envDBName)

	dsn := userName + ":" + password + "@" + "tcp(" + host + ":3306)/" + dbName + "?charset=utf8&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(db)
}
