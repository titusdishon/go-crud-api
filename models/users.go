package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/titusdishon/go-docker-mysql/config"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Summary string `json:"summary"`
}

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
}

func (u *User) CreateUser() *User {
	stmt, err := db.Prepare("INSERT INTO users (name, email, summary) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Name, u.Email, u.Summary)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", res)
	if err != nil {
		panic(err)
	}
	return u
}

func GetAllUsers() []*User {
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	var users []*User

	for result.Next() {
		var u User
		err = result.Scan(&u.ID, &u.Name, &u.Email, &u.Summary)

		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}
	return users
}
