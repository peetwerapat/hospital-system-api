package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/repository"
	"github.com/peetwerapat/hospital-system-api/pkg/myJwt"
	"golang.org/x/crypto/bcrypt"
)

type StaffUsecase struct {
	staffRepo    repository.StaffRepositoryInterface
	hospitalRepo repository.HospitalRepositoryInterface
}

func NewStaffUsecase(staffRepo repository.StaffRepositoryInterface, hospitalRepo repository.HospitalRepositoryInterface) *StaffUsecase {
	return &StaffUsecase{staffRepo, hospitalRepo}
}

func (uc *StaffUsecase) CreateStaff(staff *domain.Staff) error {
	exist, err := uc.staffRepo.GetByUsername(staff.Username)
	if err != nil {
		return fmt.Errorf("failed to check username: %w", err)
	}
	if exist != nil {
		return ErrUsernameExists
	}

	if len(staff.Password) < 6 {
		return ErrPasswordTooShort
	}

	hospital, err := uc.hospitalRepo.GetByID(staff.HospitalID)
	if err != nil || hospital == nil {
		return ErrHospitalNotFound
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	staff.Password = string(hashed)

	return uc.staffRepo.CreateStaff(staff)
}

func (uc *StaffUsecase) StaffLogin(staff *domain.Staff) (string, string, error) {
	if staff.Username == "" || staff.Password == "" || staff.HospitalID <= 0 {
		return "", "", ErrInvalidInput
	}

	staffRes, err := uc.staffRepo.GetByUsernameAndHospital(staff.Username, staff.HospitalID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get staff: %w", err)
	}
	if staffRes == nil || staffRes.HospitalID != staff.HospitalID {
		_ = bcrypt.CompareHashAndPassword([]byte("$2a$10$7EqJtq98hPqEX7fNZaFWoO"), []byte(staff.Password))
		return "", "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staffRes.Password), []byte(staff.Password)); err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessDuration := time.Hour
	if d, err := time.ParseDuration(os.Getenv("EXPIRE_JWT_TIME")); err == nil {
		accessDuration = d
	}
	refreshDuration := 24 * time.Hour
	if d, err := time.ParseDuration(os.Getenv("EXPIRE_REFRESH_TIME")); err == nil {
		refreshDuration = d
	}

	accessToken, err := myJwt.CreateToken(staffRes, accessDuration, true)
	if err != nil {
		return "", "", fmt.Errorf("failed to create access token: %w", err)
	}

	refreshToken, err := myJwt.CreateToken(staffRes, refreshDuration, false)
	if err != nil {
		return "", "", fmt.Errorf("failed to create refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}
