package controller

import (
	"api-rest/commons"
	"api-rest/models"
	"encoding/json"
	"net/http"
)

func getAll(writer http.ResponseWriter, request *http.Request) {
	persons := []models.Person{}
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&persons)
	json, _ := json.Marshal(persons)
	commons.SendResponse(writer, http.StatusOK, json)
}
