package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NYARAS/go-mux/model"
	"github.com/jinzhu/gorm"
)

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
}
