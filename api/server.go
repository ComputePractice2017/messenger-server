package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/messenger-server/model"
	"github.com/gorilla/mux"
)

func Run() {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSesson()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/persons", getAllPersonsHandler).Methods("GET")
	r.HandleFunc("/persons", newPersonHandler).Methods("POST")

	log.Println("Running the server on port 8008...")
	http.ListenAndServe(":8008", r)
}
