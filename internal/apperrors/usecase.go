package apperrors

import "net/http"

var (
	UserUsecaseGetUserByNickname = AppError{
		Message:  "The GetUserByNickname operation failed",
		Code:     "USER_USECASE_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusInternalServerError,
	}
	UserUsecaseUpdateUserErr = AppError{
		Message:  "The update user operation failed",
		Code:     "USER_USECASE_UPDATE_USER_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserUsecaseGetUserErr = AppError{
		Message:  "The get user operation failed",
		Code:     "USER_USECASE_GET_USER_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserUsecaseCheckUserByNickErr = AppError{
		Message:  "The check user by nick operation failed",
		Code:     "USER_USECASE_CHECK_USER_BY_NICK_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserUsecaseCheckUserByNickBusyErr = AppError{
		Message:  "The user's nickname is busy",
		Code:     "USER_USECASE_CHECK_USER_BY_NICK_BUSY_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserUsecaseDeleteUserErr = AppError{
		Message:  "The delete user operation failed",
		Code:     "USER_USECASE_DELETE_USER_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)
