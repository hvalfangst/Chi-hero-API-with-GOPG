package dto

import "Chi-hero-API-with-GOPG/models"

type HeroResponse struct {
	Success bool         `json:"success"`
	Error   string       `json:"error"`
	Hero    *models.Hero `json:"hero"`
}
