package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var datoStore = make(map[int]string)

	datoStore[1] = "alfa"
	datoStore[2] = "beta"
	datoStore[3] = "gama"
	datoStore[4] = "omega"
	datoStore[5] = "delta"
	datoStore[6] = "epsilon"

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(datoStore)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/control/lista", GetNoteHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8090",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
