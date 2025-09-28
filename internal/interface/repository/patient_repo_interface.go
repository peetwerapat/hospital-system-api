package repository

import "github.com/peetwerapat/hospital-system-api/internal/domain"

type PatientRepositoryInterface interface {
	GetPatientsByHospitalID(hospitalID int, filters map[string]string) ([]domain.Patient, error)
}
