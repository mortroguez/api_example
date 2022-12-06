package controllers

import (
	"encoding/json"
	"mortroguez/api_example/commons"
	"mortroguez/api_example/models"
	"net/http"
)

func GetAllPersons(writer http.ResponseWriter, request *http.Request) {
	persons := []models.Person{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persons)

	json, _ := json.Marshal(persons)

	commons.SendResponse(writer, http.StatusOK, json)
}
