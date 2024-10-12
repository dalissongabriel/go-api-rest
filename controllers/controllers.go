package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dalissongabriel/go-api-rest/database"
	"github.com/dalissongabriel/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func FindAllCelebrities(w http.ResponseWriter, r *http.Request) {
	celebrities := []models.Celebrity{}
	database.DB.Find(&celebrities)

	json.NewEncoder(w).Encode(celebrities)
}

func FindOneCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["celebrity_id"]
	var foundCelebrity *models.Celebrity

	result := database.DB.First(&foundCelebrity, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string](string){
			"message": "data not found",
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foundCelebrity)

	}
}
