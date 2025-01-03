package http

import (
	"github.com/metju-ac/train-me-maybe/internal/repositories"
)

type Handler struct {
	UserRepo repositories.UserRepository
}
