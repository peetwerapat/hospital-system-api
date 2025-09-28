package mocks

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type StaffRepo struct {
	mock.Mock
}

func (m *StaffRepo) GetByUsername(username string) (*domain.Staff, error) {
	args := m.Called(username)
	if staff, ok := args.Get(0).(*domain.Staff); ok {
		return staff, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *StaffRepo) GetByUsernameAndHospital(username string, hospitalID int) (*domain.Staff, error) {
	args := m.Called(username, hospitalID)
	if staff, ok := args.Get(0).(*domain.Staff); ok {
		return staff, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *StaffRepo) CreateStaff(staff *domain.Staff) error {
	args := m.Called(staff)
	return args.Error(0)
}
