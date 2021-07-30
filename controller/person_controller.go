package controller

import (
	"api-rest/commons"
	"api-rest/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	persons := []models.Person{}
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&persons)
	json, _ := json.Marshal(persons)
	commons.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}

	id := mux.Vars(request)["id"]
	db := commons.GetConnection()

	defer db.Close()

	db.Find(&person, id)

	if person.ID > 0 {
		json, _ := json.Marshal(person)
		commons.SendResponse(writer, http.StatusOK, json)

	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	//we  declare an object called person
	person := models.Person{}
	//we stablish conection and close it with db
	db := commons.GetConnection()
	defer db.Close()

	//catching an error and making good practices

	error := json.NewDecoder(request.Body).Decode(person)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Create(&person).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(person)

	commons.SendResponse(writer, http.StatusCreated, json)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	//we  declare an object called person
	person := models.Person{}
	//we stablish conection and close it with db
	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&person, id)

	if person.ID > 0 {
		db.Delete(person)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}
