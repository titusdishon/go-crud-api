package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/titusdishon/go-docker-mysql/entity"
	errorHandler "github.com/titusdishon/go-docker-mysql/errors"
	"github.com/titusdishon/go-docker-mysql/services"
)

var (
	userService services.UserService = services.NewUserService()
)

type IUserController interface {
	PingMe(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteAUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewUserController() IUserController {
	return &controller{}
}

func (*controller) PingMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Hit the home endpoint %s", r.URL.Path)
	_, _ = fmt.Fprintf(w, "Welcome to the home  %s", r.URL.Query().Get("userId"))
}

func (*controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := userService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error":"Error processing your request"}`))
	}
	_ = json.NewEncoder(w).Encode(users)
}

func (*controller) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid user id"))
		return
	}
	user, err := userService.FindById(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("User not found"))
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}

func (*controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Wrong data format"})
	}
	validateErr := userService.Validate(&user)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: validateErr.Error()})
	}
	result, saveErr := userService.Save(&user)
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Error saving the user"})
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)
}

func (*controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Invalid user id"})
		return
	}
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Wrong data format"})
	}

	userDetails, err := userService.FindById(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "User not found"})
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
	userService.Update(&user, ID)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

func (*controller) DeleteAUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Invalid user id"})
		return
	}
	rows, err := userService.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "Invalid user id"})
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "User does not exist"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorHandler.ServiceError{Message: "user has been deleted successfully"})
}
