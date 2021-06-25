package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/manoj-JU/rest-api-with-unit-test/utils"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	PhoneNo  string `json:"phone_number"`
}

var users []User

func init() {
	users = []User{}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UserID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if userId, ok := params["id"]; ok {
		for _, user := range users {
			if user.UserID == userId {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
	}
	utils.HandleError(errors.New("Invalid id"), http.StatusBadRequest, w)
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", AddUser).Methods("POST")
	r.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", GetUserByID).Methods("GET")
	http.ListenAndServe(":4563", r)
}
