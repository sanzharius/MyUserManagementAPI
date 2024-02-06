package apperrors

import "net/http"

var (
	UserHashGenerateFromPassword = AppError{
		Message:  "The user hash password generate operation failed",
		Code:     "USER_HASH_GENERATE_FROM_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
	UserComparePasswordsCompareHashAndPassword = AppError{
		Message:  "The user compare passwords operation failed",
		Code:     "USER_COMPARE_PASSWORDS_COMPARE_HASH_AND_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
)
