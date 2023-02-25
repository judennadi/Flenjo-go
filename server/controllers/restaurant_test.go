package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/judennadi/flenjo-go/config"
	"github.com/judennadi/flenjo-go/models"
	"github.com/judennadi/flenjo-go/routes"
)

var router = routes.InitRoutes()
var ctx = context.Background()
var redisClient = config.GetRedisCache()

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	config.ConnectDB(os.Getenv("DEV_DB_URL"))
	setup()
	code := m.Run()
	redisClient.FlushAll(ctx)
	os.Exit(code)
}

func setup() {
	dropTables()
	models.CreateUserTable()
}

func dropTables() {
	db := config.GetDB()
	_, err := db.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		fmt.Println(err)
	}
}

func runRequest(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func checkStatusCode(t *testing.T, expected, status int) {
	if status != expected {
		t.Errorf("handler returned wrong status code: got %v expected %v", status, expected)
	}
}

func TestGetRestaurants(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/restaurants?term=&categories=food,restaurants&page=0&rating=", nil)
	res := runRequest(req)

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	expected := len(data["data"].([]interface{}))
	if expected != 30 {
		t.Errorf("expected 30, got %v", expected)
	}
}

func TestGetBars(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/restaurants/bars?term=&categories=beergardens,bars&page=0&rating=", nil)
	res := runRequest(req)

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	expected := len(data["data"].([]interface{}))
	if expected != 30 {
		t.Errorf("expected 30, got %v", expected)
	}
}

func TestGetHotel(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/restaurants/hotels?term=&categories=hotels&page=0&rating=", nil)
	res := runRequest(req)

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	expected := len(data["data"].([]interface{}))
	if expected != 30 {
		t.Errorf("expected 30, got %v", expected)
	}
}

func TestSearchAutocomplete(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/restaurants/search/autocomplete?text=piz", nil)
	res := runRequest(req)

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	expected := data["terms"].([]interface{})[0].(map[string]interface{})["alias"]
	if expected != "pizza" {
		t.Errorf("expected pizza, got %v", expected)
	}
}

func TestGetBusiness(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/restaurants?term=&categories=food,restaurants&page=0&rating=", nil)
	res := runRequest(req)

	var oldData map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &oldData)

	checkStatusCode(t, http.StatusOK, res.Code)

	oldBusinessID := oldData["data"].([]interface{})[0].(map[string]interface{})["id"].(string)

	req, _ = http.NewRequest("GET", "/api/v1/restaurants/"+oldBusinessID, nil)
	res = runRequest(req)

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	checkStatusCode(t, http.StatusOK, res.Code)

	newBusinessID := data["restaurant"].(map[string]interface{})["id"].(string)

	if newBusinessID != oldBusinessID {
		t.Errorf("expected %v, got %v", oldBusinessID, newBusinessID)
	}
}
