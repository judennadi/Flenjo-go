package routes

import (
	"github.com/gorilla/mux"
	"github.com/judennadi/flenjo-go/controllers"
)

func RestaurantRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/restaurants", controllers.GetRestaurants).Methods("GET")
	router.HandleFunc("/api/v1/restaurants/bars", controllers.GetBars).Methods("GET")
	router.HandleFunc("/api/v1/restaurants/hotels", controllers.GetHotels).Methods("GET")
	router.HandleFunc("/api/v1/restaurants/search/autocomplete", controllers.SearchAutocomplete).Methods("GET")
	router.HandleFunc("/api/v1/restaurants/{id}", controllers.GetBusiness).Methods("GET")
}
