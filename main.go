package main

import (
	"api-rest/commons"
	"api-rest/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	log.Println("Servidor funcando")
	log.Println(server.ListenAndServe())
}
