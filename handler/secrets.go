package handler

import (
	"net/http"

	"github.com/NYARAS/go-mux/model"
)

func GetSecret(w http.ResponseWriter, r *http.Request) {
	accessKey, _ := GetEnvVar("AWS_ACCESS_KEY_ID")
	secretAccessKey, _ := GetEnvVar("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" && secretAccessKey == "" {
		secretResp := model.SecretResponse{
			AccessKey:       "Opps!, I hold no value",
			SecretAccessKey: "Opps! I hold no value",
		}
		respondJSON(w, http.StatusNotFound, secretResp)
		return
	}
	secretResp := model.SecretResponse{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
	}
	respondJSON(w, http.StatusOK, secretResp)

}
