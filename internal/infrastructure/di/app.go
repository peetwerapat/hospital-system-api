package di

import (
	"github.com/peetwerapat/hospital-system-api/internal/infrastructure/repository_impl"
	"github.com/peetwerapat/hospital-system-api/internal/usecase"
	"gorm.io/gorm"
)

type AppUseCase struct {
	StaffUC   *usecase.StaffUsecase
	PatientUC *usecase.PatientUsecase
}

func InitApp(dbConn *gorm.DB) *AppUseCase {
	// Hospital
	hospitalRepo := repository_impl.NewHospitalRepositoryImplement(dbConn)

	// Staff
	staffRepo := repository_impl.NewStaffRepositoryImplement(dbConn)
	staffUC := usecase.NewStaffUsecase(staffRepo, hospitalRepo)

	// Patient
	patientRepo := repository_impl.NewPatientRepositoryImplement(dbConn)
	patientUC := usecase.NewPatientUsecase(patientRepo)

	return &AppUseCase{
		StaffUC:   staffUC,
		PatientUC: patientUC,
	}
}
