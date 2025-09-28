package repository

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
)

type StaffRepositoryInterface interface {
	CreateStaff(staff *domain.Staff) error
	GetByUsername(username string) (*domain.Staff, error)
	GetByUsernameAndHospital(username string, hospitalID int) (*domain.Staff, error)
}
