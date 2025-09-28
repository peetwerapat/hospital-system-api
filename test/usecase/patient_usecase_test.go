package usecase_test

import (
	"errors"
	"testing"

	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/usecase"
	"github.com/peetwerapat/hospital-system-api/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPatientsByHospitalID_Positive(t *testing.T) {
	patientRepo := new(mocks.PatientRepo)
	uc := usecase.NewPatientUsecase(patientRepo)

	expected := []domain.Patient{{FirstNameEN: "John"}}

	patientRepo.On(
		"GetPatientsByHospitalID",
		1,
		mock.MatchedBy(func(f map[string]string) bool {
			return f["first_name"] == "John" && f["date_of_birth"] == "1990-01-01"
		}),
	).Return(expected, nil)

	filters := map[string]string{"first_name": "John", "date_of_birth": "1990-01-01"}
	patients, err := uc.GetPatientsByHospitalID(1, filters)
	assert.NoError(t, err)
	assert.Equal(t, expected, patients)

	patientRepo.AssertCalled(t, "GetPatientsByHospitalID", 1, mock.Anything)
}

func TestGetPatientsByHospitalID_Negative(t *testing.T) {
	patientRepo := new(mocks.PatientRepo)
	uc := usecase.NewPatientUsecase(patientRepo)

	t.Run("Invalid hospitalID", func(t *testing.T) {
		_, err := uc.GetPatientsByHospitalID(0, map[string]string{})
		assert.Equal(t, usecase.ErrHospitalNotFound, err)
	})

	t.Run("Filter sanitization", func(t *testing.T) {
		input := map[string]string{
			"first_name":    "Alice",
			"date_of_birth": "2023-13-01",
			"unknown":       "abc",
		}

		patientRepo.On(
			"GetPatientsByHospitalID",
			1,
			mock.MatchedBy(func(f map[string]string) bool {
				val, ok := f["first_name"]
				return ok && val == "Alice"
			}),
		).Return([]domain.Patient{}, nil)

		_, err := uc.GetPatientsByHospitalID(1, input)
		assert.NoError(t, err)
	})

	t.Run("Repo error", func(t *testing.T) {
		patientRepo.On(
			"GetPatientsByHospitalID",
			1,
			mock.Anything,
		).Return(nil, errors.New("db error"))

		_, err := uc.GetPatientsByHospitalID(1, map[string]string{"first_name": "Bob"})
		assert.ErrorContains(t, err, "db error")
	})
}
