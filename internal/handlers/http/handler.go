package http

import (
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

type Handler struct {
	Stations         []models.StationModel
	UserRepo         repositories.UserRepository
	WatchedRouteRepo repositories.WatchedRouteRepository
	ApiClient        *openapiclient.APIClient
}
