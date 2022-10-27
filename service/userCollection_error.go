package models

import (
	"net/http"

	"github.com/TendonT52/tendon-api/error"
	"github.com/gin-gonic/gin"
)

var (
	UserCollectionError = error.ErrorWithCode{
		Kind:     "collection user error",
		Code:     http.StatusInternalServerError,
		Response: gin.H{"message": "database user error"},
	}
	UserIsAlreadyExist = error.ErrorWithCode{
		Kind:     "user is already exist",
		Code:     http.StatusConflict,
		Response: gin.H{"message": "This email already in use"},
	}
	ErrorWhileAddUserToDatabase = error.ErrorWithCode{
		Kind:     "error while add user too database",
		Code:     http.StatusInsufficientStorage,
		Response: gin.H{"message": "Error while add user to database"},
	}
	EmailNotFound = error.ErrorWithCode{
		Kind:     "email not found",
		Code:     http.StatusNotFound,
		Response: gin.H{"message": "Email not found"},
	}
	IdNotFound = error.ErrorWithCode{
		Kind:     "id not found",
		Code:     http.StatusNotFound,
		Response: gin.H{"message": "Id not found"},
	}
	IdNotValid = error.ErrorWithCode{
		Kind:     "id not valid",
		Code:     http.StatusNotAcceptable,
		Response: gin.H{"message": "Id not valid"},
	}
)