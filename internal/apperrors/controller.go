package apperrors

import "net/http"

var (
	UserControllerGetUserUserNotExist = AppError{
		Message:  "The get user operation has been failed, user is not exist",
		Code:     "USER_CONTROLLER_GET_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}
	UserControllerGetUsersGetPaginationFromCtx = AppError{
		Message:  "GetUsers operation failed",
		Code:     "USER_CONTROLLER_GET_USERS_GET_PAGINATION_FROM_CTX",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerCreateUserBind = AppError{
		Message:  "The create user bind operation failed",
		Code:     "USER_CONTROLLER_CREATE_USER_BIND",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerCreateUser = AppError{
		Message:  "The create user operation failed",
		Code:     "USER_CONTROLLER_CREATE_USER_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	VerifyAuthUserGetUserByNickname = AppError{
		Message:  "The verify user operation failed",
		Code:     "VERIFY_AUTH_USER_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}
	VerifyAuthUserComparePasswords = AppError{
		Message:  "The verify user operation failed, password is incorrect",
		Code:     "VERIFY_AUTH_USER_COMPARE_PASSWORDS",
		HTTPCode: http.StatusUnauthorized,
	}
	UserBSONMarshalErr = AppError{
		Message:  "Failed to marshal user",
		Code:     "USER_BSON_MARSHAL_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerUuidParseErr = AppError{
		Message:  "Failed to parse uuid",
		Code:     "USER_CONTROLLER_UUID_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerGetUserErr = AppError{
		Message:  "The get user operation failed",
		Code:     "USER_CONTROLLER_GET_USER_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerGetUserJSON = AppError{
		Message:  "The get user operation failed",
		Code:     "USER_CONTROLLER_GET_USER_JSON",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerUpdateUserUUIDParseErr = AppError{
		Message:  "The update user failed, uuid parse error",
		Code:     "USER_USECASE_UPDATE_USER_UUID_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerUpdateUserNotExist = AppError{
		Message:  "The update user operation failed, user not exist",
		Code:     "USER_CONTROLLER_UPDATE_USER_NOT_EXIST",
		HTTPCode: http.StatusInternalServerError,
	}
	UserControllerUpdateUserBind = AppError{
		Message:  "The update user operation failed, user bind error",
		Code:     "USER_CONTROLLER_UPDATE_USER_BIND",
		HTTPCode: http.StatusInternalServerError,
	}
	ValidatorCustomValidatorValidateErr = AppError{
		Message:  "The custom validate operation failed",
		Code:     "VALIDATOR_CUSTOM_VALIDATOR_VALIDATE_ERR",
		HTTPCode: http.StatusBadRequest,
	}
	UserControllerDeleteUserUUIDParseErr = AppError{
		Message:  "The delete user operation failed, uuid parse error",
		Code:     "USER_CONTROLLER_DELETE_USER_UUID_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)
