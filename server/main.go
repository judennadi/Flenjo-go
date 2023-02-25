package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/judennadi/flenjo-go/config"
	"github.com/judennadi/flenjo-go/models"
	"github.com/judennadi/flenjo-go/routes"
)

func main() {
	godotenv.Load()

	config.ConnectDB(os.Getenv("PROD_DB_URL"))
	models.CreateUserTable()

	router := routes.InitRoutes()

	http.Handle("/", router)
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, router))
}
