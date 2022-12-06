package routes

import (
	"mortroguez/api_example/controllers"

	"github.com/gorilla/mux"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api/v1/person").Subrouter()
	subRoute.HandleFunc("/", controllers.GetAllPersons).Methods("GET")
}
