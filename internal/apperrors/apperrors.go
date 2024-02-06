package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	EnvConfigLoadErr = AppError{
		Message: "Failed to load env file",
		Code:    "ENV_INIT_ERR",
	}
	EnvConfigParseErr = AppError{
		Message: "Failed to parse env file",
		Code:    "ENV_PARSE_ERR",
	}
	MongoDBDataNotFoundErr = AppError{
		Message: "data not found in mongodb",
		Code:    "MONGO_DB_DATA_NOT_FOUND_ERR",
	}
	MongoDBFindErr = AppError{
		Message: "could not find data in mongodb",
		Code:    "MONGO_DB_FIND_ERR",
	}
	MongoDBFindOneErr = AppError{
		Message: "could not find user in mongodb",
		Code:    "MONGO_DB_FIND_ONE_ERR",
	}
	MongoDBFindOneByIDErr = AppError{
		Message: "could not find data by id in mongodb",
		Code:    "MONGO_DB_FIND_ONE_BY_ID_ERR",
	}
	MongoDBCursorErr = AppError{
		Message: "Got cursor error in mongodb",
		Code:    "MONGO_DB_CURSOR_ERR",
	}
	MongoDBUpdateErr = AppError{
		Message: "Could not update user",
		Code:    "MONGO_DB_UPDATE_ERR",
	}
	MongoDBDeleteErr = AppError{
		Message: "Could not delete user",
		Code:    "MONGO_DB_DELETE_ERR",
	}
	MongoDBDeletedCountErr = AppError{
		Message: "Could not find deleted users",
		Code:    "MONGO_DB_DELETED_COUNT_ERR",
	}
	MongoDBInsertErr = AppError{
		Message: "Cannot insert data into MongoDB",
		Code:    "MONGODB_INSERT_ERR",
	}
	MongoDBCastErr = AppError{
		Message: "Got cast error",
		Code:    "MONGODB_CAST_ERR",
	}
	NilMongoDBConfigErr = AppError{
		Message: "MongoDB config cannot be nil",
		Code:    "NIL_MONGODB_CONFIG_ERR",
	}
	MongoDBInitErr = AppError{
		Message: "Cannot init MongoDB",
		Code:    "MONGODB_INIT_ERR",
	}
	MongoDBStartErr = AppError{
		Message: "Cannot start MongoDB",
		Code:    "MONGODB_START_ERR",
	}
	ServerStartErr = AppError{
		Message:  "Failed to start app",
		Code:     "SERVER_START_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ":" + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%v: %v", appError.Message, anyErrs),
		Code:    appError.Code,
	}
}
