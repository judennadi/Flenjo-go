package main

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/judennadi/flenjo-go/config"
)

var db = config.GetDB()
var ctx = context.Background()
var redisClient = config.GetRedisCache()

func TestMain(m *testing.M) {
	godotenv.Load()

	config.ConnectDB(os.Getenv("DEV_DB_URL"))

	code := m.Run()
	redisClient.FlushAll(ctx)
	os.Exit(code)
}
