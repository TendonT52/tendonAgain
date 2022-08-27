package error

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ErrorWithCode struct {
	Kind     string
	Code     int
	Response gin.H
	Err      error
}

func (e ErrorWithCode) Error() string {
	return e.Kind
}

func (e ErrorWithCode) Unwrap() error {
	return errors.New(e.Kind)
}

func (e ErrorWithCode) With(response gin.H) *ErrorWithCode {
	e.Response = response
	return &e
}

func (e ErrorWithCode) Is(target error) bool {
	x, ok := target.(ErrorWithCode)
	if ok  {
		return x.Kind == e.Kind
	}else {
		return target.Error() == e.Error()
	}
}

func (e ErrorWithCode) From(err error) ErrorWithCode {
	e.Err = err
	return e
}

func (e ErrorWithCode) New() ErrorWithCode {
	return e
}

func (e ErrorWithCode) Dig() error {
	return e.Err
}