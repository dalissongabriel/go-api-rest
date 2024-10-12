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

func CreateCelebrity(w http.ResponseWriter, r *http.Request) {
	var newCelebrity models.Celebrity
	json.NewDecoder(r.Body).Decode(&newCelebrity)
	database.DB.Create(&newCelebrity)
	json.NewEncoder(w).Encode(newCelebrity)
}

func UpdateCelebrity(w http.ResponseWriter, r *http.Request) {
	var celebrityFromBodyRequest models.Celebrity
	json.NewDecoder(r.Body).Decode(&celebrityFromBodyRequest)

	var foundCelebrity *models.Celebrity
	resultCelebrityFirst := database.DB.First(&foundCelebrity, celebrityFromBodyRequest.Id)

	if resultCelebrityFirst.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string](string){
			"message": "data not found",
		})
		return
	}

	resultSave := database.DB.Save(&celebrityFromBodyRequest)

	if resultSave.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string](string){
			"message": "failed to updated",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(celebrityFromBodyRequest)

}

func DeleteCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["celebrity_id"]

	result := database.DB.Delete(models.Celebrity{}, id)
	fmt.Println(result.RowsAffected)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string](string){
			"message": "bad request",
		})

		return

	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string](string){
			"message": "data not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string](string){
		"message": "deleted with success",
	})
}
