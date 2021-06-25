package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {

	var jsonStr = []byte(`{
		"user_name": "Manoj",
		"phone_number": "23123123123"
		}`)

	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := User{
		UserName: "Manoj",
		PhoneNo:  "23123123123",
	}
	var output User
	_ = json.NewDecoder(rr.Body).Decode(&output)
	output.UserID = ""
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var output []User
	_ = json.NewDecoder(rr.Body).Decode(&output)
	expected := users
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	url := fmt.Sprintf("/api/v1/users/%s", "invalidID")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserByID(t *testing.T) {
	url := fmt.Sprintf("/api/v1/users/%s", users[0].UserID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var output User
	_ = json.NewDecoder(rr.Body).Decode(&output)

	// Check the response body is what we expect.
	expected := users[0]
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
