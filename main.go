package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/titusdishon/go-docker-mysql/models"
	"github.com/titusdishon/go-docker-mysql/utils"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hit the home endpoint")
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env files")
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", usersPage)
	http.HandleFunc("/user/create", createUser)
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(PORT, nil))
}
