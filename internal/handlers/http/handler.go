package http

import (
	"context"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"sync"
)

type Handler struct {
	Stations          map[int64]models.StationModel
	UserRepo          repositories.UserRepository
	WatchedRouteRepo  repositories.WatchedRouteRepository
	ApiClient         *openapiclient.APIClient
	GoroutineContexts map[uint]context.CancelFunc
	mu                sync.Mutex
	Config            config.Config
}
