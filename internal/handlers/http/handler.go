package http

import (
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
)

type Handler struct {
	Stations []models.StationModel
	UserRepo repositories.UserRepository
}
