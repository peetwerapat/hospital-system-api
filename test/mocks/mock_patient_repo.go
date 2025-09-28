package mocks

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type PatientRepo struct {
	mock.Mock
}

func (m *PatientRepo) GetPatientsByHospitalID(hospitalID int, filters map[string]string) ([]domain.Patient, error) {
	args := m.Called(hospitalID, filters)
	if patients, ok := args.Get(0).([]domain.Patient); ok {
		return patients, args.Error(1)
	}
	return nil, args.Error(1)
}
