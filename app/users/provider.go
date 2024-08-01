package users

import (
	"database/sql"

	"github.com/google/wire"
	"go.mod/app/domain"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideHandler,
	ProvideService,
	ProvideRepository,

	// bind each one of the interfaces
	wire.Bind(new(domain.UserHandler), new(*Handler)),
	wire.Bind(new(domain.UserService), new(*Service)),
	wire.Bind(new(domain.UserRepository), new(*Repository)),
)

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func ProvideService(repo domain.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func ProvideHandler(service domain.UserService) *Handler {
	return &Handler{
		svc: service,
	}
}
