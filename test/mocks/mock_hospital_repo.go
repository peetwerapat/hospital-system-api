package mocks

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type HospitalRepo struct {
	mock.Mock
}

func (m *HospitalRepo) GetByID(id int) (*domain.Hospital, error) {
	args := m.Called(id)
	if hospital, ok := args.Get(0).(*domain.Hospital); ok {
		return hospital, args.Error(1)
	}
	return nil, args.Error(1)
}
