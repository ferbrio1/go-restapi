package main

import (
	"net/http"

	"github.com/ferbrio1/ApiRest/db"
	"github.com/ferbrio1/ApiRest/models"
	"github.com/ferbrio1/ApiRest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	// Realiza migraciones de base de datos
	db.DBAutoMigrate(models.IceCream{})

	// Inicializa el enrutador
	r := mux.NewRouter()

	// Define las rutas
	r.HandleFunc("/icecreams", routes.PostIceCream).Methods("POST")
	r.HandleFunc("/icecreams", routes.GetAll).Methods("GET")
	r.HandleFunc("/icecreams/{id}", routes.GetIceHandler).Methods("GET")     // Cambiado de GetIcesHandler a GetIceHandler
	r.HandleFunc("/icecreams/{id}", routes.DeleteIceCream).Methods("DELETE") // Cambiado de GetIcesHandler a DeleteIceHandler
	r.HandleFunc("/icecreams/{id}", routes.UpdateIceCream).Methods("PUT")    // Cambiado de GetIcesHandler a UpdateIceHandler

	// Inicia el servidor
	http.ListenAndServe(":3000", r)
}
