package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB(DB_URL string) {
	database, _ := sql.Open("postgres", DB_URL)

	err := database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")

	db = database
}

func GetRedisCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.43.180:6379",
		Password: "",
		DB:       0,
	})
}

func GetDB() *sql.DB {
	return db
}
