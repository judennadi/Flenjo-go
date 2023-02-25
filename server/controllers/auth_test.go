package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/judennadi/flenjo-go/models"
)

var user = models.User{
	Name:     "Jude Nnadi",
	Username: "hazard",
	Email:    "judeonnadi@gmail.com",
	Password: "hazard",
}

func TestRegister(t *testing.T) {
	// setup()
	var resUser models.User
	var data map[string]models.User

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
	res := runRequest(req)

	json.Unmarshal(res.Body.Bytes(), &data)
	resUser = data["data"]
	checkStatusCode(t, http.StatusCreated, res.Code)

	if user.Email != resUser.Email && user.Username != resUser.Username {
		t.Errorf("expected email: %v, username: %v. got email: %v, username: %v", user.Email, user.Username, resUser.Email, resUser.Username)
	}
}

func TestLogin(t *testing.T) {
	setup()
	user.CreateUser()
	var data map[string]models.User
	var resUser models.User

	reqBody := models.User{
		Username: user.Username,
		Password: user.Password,
	}
	bodyByte, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(bodyByte))
	res := runRequest(req)
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)
	resUser = data["data"]

	if resUser.Username != reqBody.Username {
		t.Errorf("expected username: %v. got username: %v", reqBody.Username, resUser.Username)
	}
}

func TestLogout(t *testing.T) {
	var data map[string]string
	req, _ := http.NewRequest("GET", "/api/v1/auth/logout", nil)
	res := runRequest(req)
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	expected := data["user"]
	if expected != "logged out" {
		t.Errorf("expected: '%v', got: '%v'", "logged out", expected)
	}
}
