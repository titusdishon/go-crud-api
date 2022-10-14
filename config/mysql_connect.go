package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env files")
	}
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	d, err := sql.Open("mysql", DBURL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("I am connected")
	}
	db = d
}

func GetDb() *sql.DB {

	return db
}
