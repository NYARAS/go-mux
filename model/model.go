package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SecretResponse struct {
	AccessKey       string `json:"accesskey"`
	SecretAccessKey string `json:"secretaccesskey"`
}
