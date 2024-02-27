package controller

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/domain/model"
	"myAPIProject/internal/usecase/usecase"
	"myAPIProject/internal/utils"
	"net/http"
	"time"
)

const UserAuthCtx = "userAuth"

type userController struct {
	userUseCase usecase.User
}

type IUserController interface {
	GetUsers(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	BasicAuth() echo.MiddlewareFunc
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

func NewUserController(userUseCase usecase.User) IUserController {
	return &userController{userUseCase}
}

func (uc *userController) GetUsers(ctx echo.Context) error {
	var users []*model.User
	mongoDBContext, cancel := context.WithTimeout(ctx.Request().Context(), time.Second*5)
	defer cancel()

	paginationQuery, err := utils.GetPaginationFromCtx(ctx.QueryParam("page"), ctx.QueryParam("size"), ctx.QueryParam("orderBy"))
	if err != nil {
		appError := apperrors.UserControllerGetUsersGetPaginationFromCtx.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	users, err = uc.userUseCase.List(mongoDBContext, paginationQuery)
	if err != nil {
		appError := apperrors.UserControllerGetUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}

func (uc *userController) GetUser(ctx echo.Context) error {
	userUUID := ctx.Param("id")

	uid, err := uuid.Parse(userUUID)
	if err != nil {
		return apperrors.UserControllerUuidParseErr.AppendMessage(err)
	}

	user, err := uc.userUseCase.GetUser(ctx.Request().Context(), uid)
	if err != nil {
		return apperrors.UserControllerGetUserErr.AppendMessage(err)
	}

	err = ctx.JSON(http.StatusOK, user)
	if err != nil {
		return apperrors.UserControllerGetUserJSON.AppendMessage(err)
	}

	return nil
}

func (uc *userController) CreateUser(ctx echo.Context) error {
	params := &model.User{}
	//mongoDBContext, cancel := context.WithTimeout(ctx.Request().Context(), time.Second*5)
	//defer cancel()

	if err := ctx.Bind(params); err != nil {
		appError := apperrors.UserControllerCreateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	createdUser, err := uc.userUseCase.Create(ctx.Request().Context(), params)
	fmt.Printf("createdUser= %s", createdUser)
	if err != nil {
		appError := apperrors.UserControllerCreateUser.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusCreated, createdUser)
}

func (uc *userController) BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(uc.VerifyAuthUser())
}

func (uc *userController) VerifyAuthUser() func(username, password string, ctx echo.Context) (bool, error) {
	return func(username, password string, ctx echo.Context) (bool, error) {
		user, err := uc.userUseCase.GetUserByNickname(ctx.Request().Context(), username)
		if err != nil {
			return false, apperrors.VerifyAuthUserGetUserByNickname.AppendMessage(err)
		}
		if user == nil {
			return false, apperrors.VerifyAuthUserGetUserByNickname.AppendMessage(err)
		}

		err = user.ComparePasswords(password)
		if err != nil {
			return false, apperrors.VerifyAuthUserComparePasswords.AppendMessage(err)
		}

		ctx.Set(UserAuthCtx, user)

		return true, nil
	}
}

func (uc *userController) UpdateUser(ctx echo.Context) error {
	userUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		appError := apperrors.UserControllerUpdateUserUUIDParseErr.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	user, err := uc.userUseCase.GetUser(ctx.Request().Context(), userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if user == nil {
		appError := apperrors.UserControllerUpdateUserNotExist.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if err = ctx.Bind(user); err != nil {
		appError := apperrors.UserControllerUpdateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	authUser := uc.FetchAuthUser(ctx, UserAuthCtx)
	err = user.HasPermissionToUpdateUser(authUser)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	passUser := user.Password
	if passUser != user.Password {
		err = user.HashPassword()
	}

	err = ctx.Validate(user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	_, err = uc.userUseCase.CheckUserByNickname(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	updatedUser, err := uc.userUseCase.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *userController) FetchAuthUser(ctx echo.Context, UserAuthCtx string) *model.User {
	return ctx.Get(UserAuthCtx).(*model.User)
}

func (uc *userController) DeleteUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		appError := apperrors.UserControllerDeleteUserUUIDParseErr.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	err = uc.userUseCase.DeleteUser(ctx.Request().Context(), &userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, userUUID)
}
