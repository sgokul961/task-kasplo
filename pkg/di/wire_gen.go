// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"main.go/pkg/api"
	"main.go/pkg/api/handler"
	"main.go/pkg/config"
	"main.go/pkg/db"
	"main.go/pkg/helper"
	"main.go/pkg/repo"
	"main.go/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*apis.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	loginrepo := repo.NewLoginRepo(gormDB)
	interfaceshelperHelper := helper.NewHelper(cfg)
	lOginUseCase := usecase.NewLoginUseCase(loginrepo, interfaceshelperHelper)
	loginHNadler := handler.NewLoginHandler(lOginUseCase)
	serverHTTP := apis.NewServerHttp(loginHNadler)
	return serverHTTP, nil
}
