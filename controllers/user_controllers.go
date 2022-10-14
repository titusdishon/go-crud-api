package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/titusdishon/go-docker-mysql/models"
	"github.com/titusdishon/go-docker-mysql/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hit the home endpoint")
}

func UsersPage(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}
