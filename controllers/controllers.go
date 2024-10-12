package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dalissongabriel/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func FindAllCelebrities(w http.ResponseWriter, r *http.Request) {
	celebrities := []models.Celebrity{
		{Id: "c87b56e5-bcd4-4512-b207-4ea7e2b40a28", Name: "Alisson Gabriel", Age: 25},
		{Id: "4401f966-21f3-4842-82d0-9adee409e259", Name: "Ana Costa", Age: 24},
	}
	json.NewEncoder(w).Encode(celebrities)
}

func FindOneCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	celebrities := []models.Celebrity{
		{Id: "c87b56e5-bcd4-4512-b207-4ea7e2b40a28", Name: "Alisson Gabriel", Age: 25},
		{Id: "4401f966-21f3-4842-82d0-9adee409e259", Name: "Ana Costa", Age: 24},
	}

	celebrityIdFromRequest := vars["celebrity_id"]
	var foundCelebrity *models.Celebrity

	for _, celebrity := range celebrities {
		if celebrity.Id == celebrityIdFromRequest {
			foundCelebrity = &celebrity
			break
		}
	}

	if foundCelebrity != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&foundCelebrity)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
