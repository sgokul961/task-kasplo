//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	apis "main.go/pkg/api"
	"main.go/pkg/api/handler"
	"main.go/pkg/config"
	"main.go/pkg/db"
	"main.go/pkg/helper"
	"main.go/pkg/repo"
	"main.go/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*apis.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repo.NewLoginRepo,
		usecase.NewLoginUseCase,
		handler.NewLoginHandler,
		apis.NewServerHttp,
		helper.NewHelper,
	)

	return nil, nil
}

//mistyped servmux and manually changed to serverHTTP
