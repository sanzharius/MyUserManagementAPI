package registry

import (
	"myAPIProject/internal/adapter/controller"
	"myAPIProject/internal/adapter/repository"
	"myAPIProject/internal/usecase/usecase"
)

func (r *registry) NewUserController() controller.IUserController {
	userUsecase := usecase.NewUserUsecase(
		repository.NewUserRepository(r.collection, r.db, r.logger),
		/*repository.NewDBRepository(r.collection, r.db, r.logger)*/)

	return controller.NewUserController(userUsecase)
}
