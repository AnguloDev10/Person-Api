package routes

import (
	"api-rest/controller"

	"github.com/gorilla/mux"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/person/api-v1").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	//to-do: Continuos with other routes crud

}
