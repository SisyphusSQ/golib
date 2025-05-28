package utils

import (
	"errors"
	"net/http"
)

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	InternalERROR         = 500
	InvalidParams         = 400
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Not Found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Param Invalid")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	case errors.Is(err, ErrInternalServerError):
		return http.StatusInternalServerError
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
