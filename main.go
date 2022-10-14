package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hit the home endpoint")
}

func getUsers() []*User {

	//open db connection
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//execute query
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	var users []*User

	for result.Next() {
		var u User
		err = result.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}
	return users

}
func usersPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	fmt.Println(w, "Welcome to the users page")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", usersPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
