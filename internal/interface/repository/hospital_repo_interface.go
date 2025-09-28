package repository

import "github.com/peetwerapat/hospital-system-api/internal/domain"

type HospitalRepositoryInterface interface {
	GetByID(hospitalID int) (*domain.Hospital, error)
}
