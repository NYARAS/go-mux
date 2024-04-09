package handler

import (
	"net/http"

	"encoding/json"

	"github.com/NYARAS/go-mux/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

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

func GetSecret(w http.ResponseWriter, r *http.Request) {
	accessKey, _ := GetEnvVar("AWS_ACCESS_KEY_ID")
	secretAccessKey, _ := GetEnvVar("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" && secretAccessKey == "" {
		logrus.Error("How unfortunate. I hold no secrets ")
		secretResp := model.SecretResponse{
			AccessKey:       "I hold no secrets",
			SecretAccessKey: map[string]string{},
		}
		respondJSON(w, http.StatusNotFound, secretResp)
		return
	}
	logrus.Info("Whats this?")
	maps := map[string]string{}
	maps["access_key"] = accessKey
	maps["secret_access_key"] = secretAccessKey
	secretResp := model.SecretResponse{
		AccessKey:       "Egads!! I have some secrets, I must inform you at once",
		SecretAccessKey: maps,
	}
	respondJSON(w, http.StatusOK, secretResp)

}
