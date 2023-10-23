package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ferbrio1/go-restapi/db"
	"github.com/ferbrio1/go-restapi/models"
	"github.com/gorilla1/mux"
)

// CREATE
func PostIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	json.NewDecoder(r.Body).Decode(&ice)

	createdIce := db.DB.Create(&ice)
	err := createdIce.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ice)
}

// all
func GetAll(w http.ResponseWriter, r *http.Request) {
	var ices []models.IceCream
	db.DB.Find(&ices)

	json.NewEncoder(w).Encode(&ices)
}

// Leer
func GetIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}
	json.NewEncoder(w).Encode(&ice)
}

// Actualizar
func UpdateIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	var updatedIce models.IceCream

	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}

	json.NewDecoder(r.Body).Decode(&updatedIce)
	db.DB.Model(&ice).Updates(updatedIce)

	json.NewEncoder(w).Encode(&ice)
}

// Borrar
func DeleteIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}

	db.DB.Unscoped().Delete(&ice)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&ice)
}
