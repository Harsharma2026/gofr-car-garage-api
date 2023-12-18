package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Harsharma2026/car-garage-api/db"
	"github.com/Harsharma2026/car-garage-api/models"
	"github.com/gorilla/mux"
)

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := db.GetDatabase()
	err = db.Create(&car).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetDatabase()
	var cars []models.Car
	err := db.Find(&cars).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func GetCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := db.GetDatabase()
	var car models.Car
	err = db.First(&car, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateCar models.Car
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updateCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := db.GetDatabase()
	err = db.Model(&models.Car{ID: id}).Updates(updateCar).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateCar)
}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := db.GetDatabase()
	err = db.Delete(&models.Car{ID: id}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
