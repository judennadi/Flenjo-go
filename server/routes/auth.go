package routes

import (
	"github.com/gorilla/mux"
	"github.com/judennadi/flenjo-go/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	AuthRoutes(router)
	RestaurantRoutes(router)
	return router
}

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/logout", controllers.Logout).Methods("GET")
}
