package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/titusdishon/go-docker-mysql/models"
	"github.com/titusdishon/go-docker-mysql/utils"
)

func PingMe(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hit the home endpoint %s", r.URL.Path)
	fmt.Fprintf(w, "Welcome to the home  %s", r.URL.Query().Get("userId"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user id"))
		return
	}
	user := models.GetUserById(ID)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func DeleteAUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user id"))
		return
	}
	rows := models.DeleteUser(ID)
	if rows != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User does not exist"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted successfully"))
}
