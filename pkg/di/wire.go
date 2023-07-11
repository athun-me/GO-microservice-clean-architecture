//go:build wireinject
// +build wireinject

package di

import (
	http "athun/pkg/api"
	handler "athun/pkg/api/handler"
	config "athun/pkg/config"
	db "athun/pkg/db"
	repo "athun/pkg/repository"
	useCase "athun/pkg/usecase"

	"github.com/google/wire"
)

func InitApi(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repo.NewUserRepo,
		useCase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHttp,
	)
	return &http.ServerHttp{}, nil

}

//go run github.com/google/wire/cmd/wire
