package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Summary string `json:"summary"`
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
		err = result.Scan(&u.ID, &u.Name, &u.Email, &u.Summary)

		if err != nil {
			panic(err)
		}
		fmt.Println("Users------------------", &u)

		users = append(users, &u)
	}
	return users

}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	fmt.Println(w, users)

	json.NewEncoder(w).Encode(users)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("failed to load env files")
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", usersPage)
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(PORT, nil))
}
