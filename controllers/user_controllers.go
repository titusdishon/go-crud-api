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
	"github.com/titusdishon/go-docker-mysql/utils"
)

var (
	userService services.UserService
)

type IUserController interface {
	PingMe(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	SignIn(rw http.ResponseWriter, r *http.Request)
	DeleteAUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewUserController(service services.UserService) IUserController {
	// dependency injection
	userService = service
	return &controller{}
}

func (*controller) PingMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
func (*controller) SignIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errorHandler.ServiceError{Message: "Wrong data format"})
	}
	// validate the request first.
	if user.Email == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing"))
		return
	}

	// let’s see if the user exists
	result, _ := userService.UserExists(&user)

	if result == nil {
		// this means either the user does not exist
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("User Does not Exist"))
		return
	}

	valid := utils.CheckPasswordHash(user.Password, result.Password)
	if !valid {
		// this means the password is wrong
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Password"))
		return
	}
	tokenString, err := utils.GetSignedToken()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
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
