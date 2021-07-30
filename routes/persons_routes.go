package routes

import (
	"api-rest/controller"

	"github.com/gorilla/mux"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/person/api-v1").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controller.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")
}
