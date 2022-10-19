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

func GetUserById(id int64) *User {
	var user User
	err := db.QueryRow(`SELECT id, email, name, summary FROM users WHERE id=?;`, id).Scan(&user.ID,
		&user.Email,
		&user.Name,
		&user.Summary)

	if err != nil {
		fmt.Println(err)
	}
	return &user
}

func (u *User) UpdateUser(id int64) *User {
	stmt, err := db.Prepare(`UPDATE users SET name = ?, email = ?, summary = ? WHERE id = ?;`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Email, u.Summary, id)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func DeleteUser(id int64) int {
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		panic(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return int(rows)
}
