//go:build wireinject

package api

import (
	"go_confess_space-project/api/controller"
	"go_confess_space-project/api/repository"
	"go_confess_space-project/api/service"
	"go_confess_space-project/config"

	"github.com/google/wire"
)

func SpaceInjector() *controller.SpaceController {
	wire.Build(
		controller.NewSpaceAuthController,
		service.NewSpaceServiceImpl,
		repository.NewSpaceRepositoryImpl,
		config.Validator,
		config.DatabaseConnection,
	)
	return nil
}

func UserSpaceInjector() *controller.UserSpaceController {
	wire.Build(
		controller.NewUserSpaceController,
		service.NewUserSpaceServiceImpl,
		repository.NewUserSpaceRepositoryImpl,
		config.Validator,
		config.DatabaseConnection,
	)
	return nil
}

func MessageInjector() *service.MessageService {
	wire.Build(
		service.NewMessageService,
		repository.NewMessageRepositoryImpl,
		config.Validator,
		config.DatabaseConnection,
	)
	return nil
}

func UserSpaceLastSeenInjector() *controller.UserSpaceLastSeenController {
	wire.Build(
		controller.NewUserSpaceLastSeenController,
		service.NewUserSpaceLastSeenServiceImpl,
		repository.NewUserSpaceLastSeenRepositoryImpl,
		config.Validator,
		config.DatabaseConnection,
	)
	return nil
}

func UserSpaceLastSeenInjectorWithAuth() *controller.UserSpaceLastSeenController {
	wire.Build(
		controller.NewUserSpaceLastSeenController,
		service.NewUserSpaceLastSeenServiceImpl,
		repository.NewUserSpaceLastSeenRepositoryImpl,
		config.Validator,
		config.DatabaseConnection,
	)
	return nil
}
