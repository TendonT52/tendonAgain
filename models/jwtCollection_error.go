package models

import (
	"net/http"

	"github.com/TendonT52/tendon-api/error"
	"github.com/gin-gonic/gin"
)

var (
	jwtCollectionError = error.ErrorWithCode{
		Kind:     "collection jwt error",
		Code:     http.StatusInternalServerError,
		Response: gin.H{"message": "database jwt error"},
	}
	ErrorWhileAddJwtToDatabase = error.ErrorWithCode{
		Kind:     "can't add jwt to database",
		Code:     http.StatusInternalServerError,
		Response: gin.H{"message": "can't add jwt tot database"},
	}
	ErrorWhileUpdateJwtToDatabase = error.ErrorWithCode{
		Kind :  "Can't update jwt in database",
		Code : http.StatusInsufficientStorage,
		Response: gin.H{"message": "can't update jwt in database"},
	}
)