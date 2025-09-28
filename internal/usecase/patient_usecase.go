package usecase

import (
	"regexp"

	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/repository"
)

type PatientUsecase struct {
	patientRepo repository.PatientRepositoryInterface
}

func NewPatientUsecase(patientRepo repository.PatientRepositoryInterface) *PatientUsecase {
	return &PatientUsecase{patientRepo}
}

func (uc *PatientUsecase) GetPatientsByHospitalID(hospitalID int, filters map[string]string) ([]domain.Patient, error) {
	if hospitalID <= 0 {
		return nil, ErrHospitalNotFound
	}

	validFields := map[string]bool{
		"national_id": true, "passport_id": true,
		"first_name": true, "middle_name": true, "last_name": true,
		"date_of_birth": true, "phone_number": true, "email": true,
	}

	cleanFilters := map[string]string{}
	for k, v := range filters {
		if !validFields[k] || v == "" {
			continue
		}
		if k == "date_of_birth" {
			match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, v)
			if !match {
				continue
			}
		}
		cleanFilters[k] = v
	}

	return uc.patientRepo.GetPatientsByHospitalID(hospitalID, cleanFilters)
}
