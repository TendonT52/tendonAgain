package services

import (
	"net/http"

	"github.com/TendonT52/tendon-api/error"
	"github.com/gin-gonic/gin"
)

var (
	JwtError = error.ErrorWithCode{
		Kind:     "jwt error",
		Code:     http.StatusBadRequest,
		Response: gin.H{"message": "jwt error"},
	}
	UnverfiableToken = error.ErrorWithCode{
		Kind:     "unverfiable token",
		Code:     http.StatusBadRequest,
		Response: gin.H{"message": "Unverfiable Token"},
	}
	TokenExpired = error.ErrorWithCode{
		Kind:     "token expired",
		Code:     http.StatusForbidden,
		Response: gin.H{"message": "Token expired"},
	}
	TokenNotValidYet = error.ErrorWithCode{
		Kind:     "token not valid yet",
		Code:     http.StatusForbidden,
		Response: gin.H{"message": "Token not valid yet"},
	}
)
