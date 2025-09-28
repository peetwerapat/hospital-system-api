package repository_impl

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/repository"
	"gorm.io/gorm"
)

type HospitalRepositoryImplement struct {
	db *gorm.DB
}

func NewHospitalRepositoryImplement(db *gorm.DB) repository.HospitalRepositoryInterface {
	return &HospitalRepositoryImplement{db}
}

func (r *HospitalRepositoryImplement) GetByID(hospitalID int) (*domain.Hospital, error) {
	var hospital domain.Hospital

	if err := r.db.First(&hospital, hospitalID).Error; err != nil {
		return nil, err
	}

	return &hospital, nil
}
