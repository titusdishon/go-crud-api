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
	var (
		mysqlUser     = os.Getenv("MYSQL_USER")
		mysqlPassword = os.Getenv("MYSQL_PASSWORD")
		mysqlHost     = os.Getenv("MYSQL_HOST")
		mysqlPort     = os.Getenv("MYSQL_PORT")
		mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	)
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
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
