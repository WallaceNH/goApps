package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type Usuario struct {
	Nombre		string `json:"nombre"`
	Apellido 	string `json:"apellido"`
	Usuario 	string `json:"usuario"`
	Clave 		string `json:"clave"`
	Email    	string `json:email"`
}

var userStore = []Usuario{}

// HTTP - GET Obtener Usuario
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
	usuario, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usuario)
}

// HTTP - POST Agregar Usuario
func crearUsuario(w http.ResponseWriter, r *http.Request) {

	var userx Usuario
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&userx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Validate the User entity
	err = validate(userx)
	userx.Clave = "********"
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert User entity into User Store
	userStore = append(userStore, userx)
	w.WriteHeader(http.StatusCreated)
}

// Validate User entity
func validate(user Usuario) error {
	for _, u := range userStore {
		if u.Usuario == user.Usuario {
			return errors.New("Usuario ya existe")
		}
	}
	return nil
}
func Rutas() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", crearUsuario).Methods("POST")
	r.HandleFunc("/users", obtenerUsuario).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":9080", Rutas())
}