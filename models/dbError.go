package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errKind int

const (
	_ errKind = iota
	dbError
	userIsAlreadyExist
	errorWhileAddUserTooDatabase
	emailNotFound
	idNotFound
	idNotValid
)

var (
	DbError = DBError{
		Kind:  dbError,
		Code:  http.StatusInternalServerError,
		Value: gin.H{"message": "server error"},
	}
	UserIsAlreadyExist = DBError{
		Kind:  userIsAlreadyExist,
		Code:  http.StatusConflict,
		Value: gin.H{"message": "This email already in use"},
	}
	ErrorWhileAddUserToDatabase = DBError{
		Kind:  errorWhileAddUserTooDatabase,
		Code:  http.StatusInsufficientStorage,
		Value: gin.H{"message": "Error while add user to database"},
	}
	EmailNotFound = DBError{
		Kind:  emailNotFound,
		Code:  http.StatusNotFound,
		Value: gin.H{"message": "Email not found"},
	}
	IdNotFound = DBError{
		Kind:  idNotFound,
		Code:  http.StatusNotFound,
		Value: gin.H{"message": "Id not found"},
	}
	IdNotValid = DBError{
		Kind: idNotValid,
		Code: http.StatusNotAcceptable,
		Value: gin.H{"message": "Id not valid"},
	}
)

type errorWithCode interface {
	GetCode() int
	GetValue() gin.H
	GetKind() errKind
	Error() string
}

type DBError struct {
	Kind  errKind
	Code  int
	Value gin.H
	Err   error
}

func (e DBError) Error() string {
	switch e.Kind {
	case userIsAlreadyExist:
		return "This email already in use"
	default:
		return "Error from database"
	}
}

func (e *DBError) GetCode() int {
	return e.Code
}

func (e *DBError) GetKind() errKind {
	return e.Kind
}

func (e *DBError) GetValue() gin.H {
	return e.Value
}

func (e DBError) New() *DBError {
	return &e
}

func (e DBError) With(val gin.H) *DBError {
	e.Value = val
	return &e
}

func (e DBError) From(err error) *DBError {
	e.Err = err
	return &e
}
