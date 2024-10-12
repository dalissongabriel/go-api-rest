package routes

import (
	"log"
	"net/http"

	"github.com/dalissongabriel/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/api/celebrities", controllers.FindAllCelebrities).Methods("GET")
	r.HandleFunc("/api/celebrities/{celebrity_id}", controllers.FindOneCelebrity).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
