package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahmoud24598salah/Gorm-Using-gorilla-mux.git/db"
	"gorm.io/gorm"
)

type handlre struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handlre {
	return handlre{DB: db}
}

func (h handlre) allUser(w http.ResponseWriter, r *http.Request) {
	var users []db.User
	result := h.DB.Find(&users)
	if result.Error != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func (h handlre) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var user db.User
	result := h.DB.Where("name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Println("Not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.DB.Delete(&user)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h handlre) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user db.User
	result := h.DB.Where("name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Println("Not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.Email = email
	h.DB.Save(&user)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h handlre) newUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	result := h.DB.Create(&db.User{Name: name, Email: email})
	if result.Error != nil {
		fmt.Println("failed to create a new user")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Created")
}
