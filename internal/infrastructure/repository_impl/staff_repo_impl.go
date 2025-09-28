package repository_impl

import (
	"errors"

	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/repository"
	"gorm.io/gorm"
)

type StaffRepositoryImplement struct {
	db *gorm.DB
}

func NewStaffRepositoryImplement(db *gorm.DB) repository.StaffRepositoryInterface {
	return &StaffRepositoryImplement{db}
}

func (r *StaffRepositoryImplement) CreateStaff(staff *domain.Staff) error {
	return r.db.Create(staff).Error
}

func (r *StaffRepositoryImplement) GetByUsername(username string) (*domain.Staff, error) {
	var staff domain.Staff

	if err := r.db.Where("username = ?", username).First(&staff).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}

func (r *StaffRepositoryImplement) GetByUsernameAndHospital(username string, hospitalID int) (*domain.Staff, error) {
	var staff domain.Staff

	if err := r.db.Where("username = ? AND hospital_id = ?", username, hospitalID).
		First(&staff).Error; err != nil {
		return nil, err
	}

	return &staff, nil
}
