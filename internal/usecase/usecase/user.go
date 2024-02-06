package usecase

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/domain/model"
	"myAPIProject/internal/usecase/repository"
	"myAPIProject/internal/utils"
)

type userUsecase struct {
	userRepository repository.UserRepository
	dbRepository   repository.DBRepository
}

type User interface {
	List(ctx context.Context, paginationQuery *utils.PaginationQuery) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	CheckUserByNickname(ctx context.Context, user *model.User) (bool, error)
	UpdateUser(ctx context.Context, user *model.User) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, userID *uuid.UUID) error
}

func NewUserUsecase(userRepository repository.UserRepository, dbRepository repository.DBRepository) User {
	return &userUsecase{
		userRepository: userRepository,
		dbRepository:   dbRepository,
	}
}

func (userUsecase *userUsecase) List(ctx context.Context, paginationQuery *utils.PaginationQuery) ([]*model.User, error) {
	user, err := userUsecase.userRepository.FindAll(ctx, paginationQuery)
	if err != nil {
		return nil, apperrors.MongoDBFindErr.AppendMessage(err)
	}

	return user, nil
}

func (userUsecase *userUsecase) Create(ctx context.Context, user *model.User) (*model.User, error) {
	data, err := userUsecase.dbRepository.Transaction(func(i interface{}) (interface{}, error) {
		user, err := userUsecase.userRepository.Create(ctx, user)
		if err != nil {
			return nil, apperrors.MongoDBInsertErr.AppendMessage(err)
		}

		return user, err
	})
	userInst, ok := data.(*model.User)
	if !ok {
		return nil, apperrors.MongoDBCastErr.AppendMessage(err)
	}

	if err != nil {
		return nil, apperrors.MongoDBInsertErr.AppendMessage(err)
	}

	return userInst, nil
}

func (userUsecase *userUsecase) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	user, err := userUsecase.userRepository.FindUserByNickname(ctx, nickname)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByNickname.AppendMessage(err)
	}

	return user, nil
}

func (userUsecase *userUsecase) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := userUsecase.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserErr.AppendMessage(err)
	}

	return user, nil
}

func (userUsecase *userUsecase) CheckUserByNickname(ctx context.Context, user *model.User) (bool, error) {
	checkedUser, err := userUsecase.userRepository.FindUserByNickname(ctx, user.Nickname)
	if err != nil {
		return false, apperrors.UserUsecaseCheckUserByNickErr.AppendMessage(err)
	}
	if checkedUser.ID != user.ID {
		return false, apperrors.UserUsecaseCheckUserByNickBusyErr.AppendMessage(err)
	}

	return true, nil
}

func (userUsecase *userUsecase) UpdateUser(ctx context.Context, user *model.User) (*mongo.UpdateResult, error) {
	updateResult, err := userUsecase.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, apperrors.UserUsecaseUpdateUserErr.AppendMessage(err)
	}

	return updateResult, nil
}

func (userUsecase *userUsecase) DeleteUser(ctx context.Context, userID *uuid.UUID) error {
	err := userUsecase.userRepository.DeleteOne(ctx, userID)
	if err != nil {
		return apperrors.UserUsecaseDeleteUserErr.AppendMessage(err)
	}

	return nil
}
