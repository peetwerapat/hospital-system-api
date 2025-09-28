package usecase_test

import (
	"errors"
	"os"
	"testing"

	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/usecase"
	"github.com/peetwerapat/hospital-system-api/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateStaff_Positive(t *testing.T) {
	staffRepo := new(mocks.StaffRepo)
	hospitalRepo := new(mocks.HospitalRepo)
	uc := usecase.NewStaffUsecase(staffRepo, hospitalRepo)

	staff := &domain.Staff{Username: "john", Password: "123456", HospitalID: 1}

	staffRepo.On("GetByUsername", "john").Return(nil, nil)
	hospitalRepo.On("GetByID", 1).Return(&domain.Hospital{ID: 1}, nil)
	staffRepo.On("CreateStaff", mock.Anything).Return(nil)

	err := uc.CreateStaff(staff)
	assert.NoError(t, err)
	staffRepo.AssertCalled(t, "CreateStaff", mock.Anything)
}

func TestCreateStaff_Negative(t *testing.T) {
	staffRepo := new(mocks.StaffRepo)
	hospitalRepo := new(mocks.HospitalRepo)
	uc := usecase.NewStaffUsecase(staffRepo, hospitalRepo)

	t.Run("Username exists", func(t *testing.T) {
		staffRepo.On("GetByUsername", "john").Return(&domain.Staff{Username: "john"}, nil)
		err := uc.CreateStaff(&domain.Staff{Username: "john", Password: "123456", HospitalID: 1})
		assert.Equal(t, usecase.ErrUsernameExists, err)
	})

	t.Run("Password too short", func(t *testing.T) {
		staffRepo.On("GetByUsername", "alice").Return(nil, nil)
		err := uc.CreateStaff(&domain.Staff{Username: "alice", Password: "123", HospitalID: 1})
		assert.Equal(t, usecase.ErrPasswordTooShort, err)
	})

	t.Run("Hospital not found", func(t *testing.T) {
		staffRepo.On("GetByUsername", "bob").Return(nil, nil)
		hospitalRepo.On("GetByID", 99).Return(nil, errors.New("not found"))
		err := uc.CreateStaff(&domain.Staff{Username: "bob", Password: "123456", HospitalID: 99})
		assert.Equal(t, usecase.ErrHospitalNotFound, err)
	})

	t.Run("Repo error", func(t *testing.T) {
		staffRepo.On("GetByUsername", "eve").Return(nil, errors.New("db error"))
		err := uc.CreateStaff(&domain.Staff{Username: "eve", Password: "123456", HospitalID: 1})
		assert.ErrorContains(t, err, "failed to check username")
	})
}

func TestStaffLogin_Positive(t *testing.T) {
	staffRepo := new(mocks.StaffRepo)
	hospitalRepo := new(mocks.HospitalRepo)
	uc := usecase.NewStaffUsecase(staffRepo, hospitalRepo)

	password := "123456"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	staffRepo.On("GetByUsernameAndHospital", "john", 1).Return(&domain.Staff{
		Username:   "john",
		Password:   string(hashed),
		HospitalID: 1,
	}, nil)

	staff := &domain.Staff{Username: "john", Password: password, HospitalID: 1}
	os.Setenv("EXPIRE_JWT_TIME", "1h")
	os.Setenv("EXPIRE_REFRESH_TIME", "24h")
	access, refresh, err := uc.StaffLogin(staff)
	assert.NoError(t, err)
	assert.NotEmpty(t, access)
	assert.NotEmpty(t, refresh)
}

func TestStaffLogin_Negative(t *testing.T) {
	staffRepo := new(mocks.StaffRepo)
	hospitalRepo := new(mocks.HospitalRepo)
	uc := usecase.NewStaffUsecase(staffRepo, hospitalRepo)

	t.Run("Invalid input", func(t *testing.T) {
		_, _, err := uc.StaffLogin(&domain.Staff{Username: "", Password: "123", HospitalID: 0})
		assert.Equal(t, usecase.ErrInvalidInput, err)
	})

	t.Run("Staff not found", func(t *testing.T) {
		staffRepo.On("GetByUsernameAndHospital", "notexist", 1).Return(nil, nil)
		_, _, err := uc.StaffLogin(&domain.Staff{Username: "notexist", Password: "123456", HospitalID: 1})
		assert.Equal(t, usecase.ErrInvalidCredentials, err)
	})

	t.Run("Password mismatch", func(t *testing.T) {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)
		staffRepo.On("GetByUsernameAndHospital", "john", 1).Return(&domain.Staff{Username: "john", Password: string(hashed), HospitalID: 1}, nil)
		_, _, err := uc.StaffLogin(&domain.Staff{Username: "john", Password: "wrong", HospitalID: 1})
		assert.Equal(t, usecase.ErrInvalidCredentials, err)
	})

	t.Run("Repo error", func(t *testing.T) {
		staffRepo.On("GetByUsernameAndHospital", "error", 1).Return(nil, errors.New("db error"))
		_, _, err := uc.StaffLogin(&domain.Staff{Username: "error", Password: "123456", HospitalID: 1})
		assert.ErrorContains(t, err, "failed to get staff")
	})

	t.Run("Hospital mismatch", func(t *testing.T) {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		staffRepo.On("GetByUsernameAndHospital", "john", 1).Return(&domain.Staff{Username: "john", Password: string(hashed), HospitalID: 2}, nil)
		_, _, err := uc.StaffLogin(&domain.Staff{Username: "john", Password: "123456", HospitalID: 1})
		assert.Equal(t, usecase.ErrInvalidCredentials, err)
	})
}
