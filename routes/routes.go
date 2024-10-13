package routes

import (
	"log"
	"net/http"

	"github.com/dalissongabriel/go-api-rest/controllers"
	"github.com/dalissongabriel/go-api-rest/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.LogHttpRequestURI)
	r.Use(middlewares.SetContentTypeJSON)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/api/celebrities", controllers.CreateCelebrity).Methods("POST")
	r.HandleFunc("/api/celebrities", controllers.UpdateCelebrity).Methods("PUT")
	r.HandleFunc("/api/celebrities", controllers.FindAllCelebrities).Methods("GET")
	r.HandleFunc("/api/celebrities/{celebrity_id}", controllers.FindOneCelebrity).Methods("GET")
	r.HandleFunc("/api/celebrities/{celebrity_id}", controllers.DeleteCelebrity).Methods("DELETE")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
