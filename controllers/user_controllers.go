package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/titusdishon/go-docker-mysql/entity"
	"github.com/titusdishon/go-docker-mysql/repositories"
)

var (
	repo repositories.UserRepository = repositories.NewUserRepository()
)

func PingMe(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hit the home endpoint %s", r.URL.Path)
	_, _ = fmt.Fprintf(w, "Welcome to the home  %s", r.URL.Query().Get("userId"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error":"Error processing your request"}`))
	}
	_ = json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid user id"))
		return
	}
	user, err := repo.FindById(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("User not found"))
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Error":"Wrong data format"}`))
	}
	repo.Save(&user)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid user id"))
		return
	}
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Error":"Wrong data format"}`))
	}

	userDetails, err := repo.FindById(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("User not found"))
		return
	}
	if user.Name != "" {
		userDetails.Name = user.Name
	}
	if user.Email != "" {
		userDetails.Email = user.Email
	}

	if user.Summary != "" {
		userDetails.Summary = user.Summary
	}
	repo.Update(&user, ID)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

func DeleteAUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid user id"))
		return
	}
	rows, err := repo.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid user id"))
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("User does not exist"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Deleted successfully"))
}
