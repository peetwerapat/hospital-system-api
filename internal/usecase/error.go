package usecase

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidInput         = errors.New("invalid input")
	ErrInvalidCredentials   = errors.New("invalid username or password")
	ErrPasswordTooShort     = errors.New("password must be at least 6 characters")
	ErrHospitalNotFound     = errors.New("hospital not found")
	ErrUsernameExists       = errors.New("username already exists")
	ErrInternal             = errors.New("internal error")
	ErrInvalidHospitalToken = errors.New("Invalid hospital_id in token")
	ErrInvalidQuery         = errors.New("Invalid query parameters")
	ErrInvalidHospitalID    = errors.New("hospital ID is invalid or missing")
)

func MapCreateStaffError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrUsernameExists):
		return http.StatusConflict, ErrUsernameExists.Error()
	case errors.Is(err, ErrHospitalNotFound):
		return http.StatusNotFound, ErrHospitalNotFound.Error()
	case errors.Is(err, ErrPasswordTooShort):
		return http.StatusBadRequest, ErrPasswordTooShort.Error()
	default:
		return http.StatusInternalServerError, ErrInternal.Error()
	}
}

func MapStaffLoginError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest, ErrInvalidInput.Error()
	case errors.Is(err, ErrInvalidCredentials):
		return http.StatusUnauthorized, ErrInvalidCredentials.Error()
	default:
		return http.StatusInternalServerError, ErrInternal.Error()
	}
}

func MapGetPatientsByHospitalIDError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrInvalidHospitalID):
		return http.StatusBadRequest, ErrInvalidHospitalID.Error()
	default:
		return http.StatusInternalServerError, ErrInternal.Error()
	}

}
