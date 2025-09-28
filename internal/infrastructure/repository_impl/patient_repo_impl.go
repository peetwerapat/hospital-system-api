package repository_impl

import (
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/repository"
	"gorm.io/gorm"
)

type PatientRepositoryImplement struct {
	db *gorm.DB
}

func NewPatientRepositoryImplement(db *gorm.DB) repository.PatientRepositoryInterface {
	return &PatientRepositoryImplement{db}
}

func (r *PatientRepositoryImplement) GetPatientsByHospitalID(hospitalID int, filters map[string]string) ([]domain.Patient, error) {
	var patients []domain.Patient
	db := r.db.Model(&domain.Patient{}).Where("hospital_id = ?", hospitalID)

	if id := filters["id"]; id != "" {
		db = db.Where("national_id = ? OR passport_id = ?", id, id)
	}
	if f := filters["first_name"]; f != "" {
		db = db.Where("first_name_th ILIKE ? OR first_name_en ILIKE ?", "%"+f+"%", "%"+f+"%")
	}
	if m := filters["middle_name"]; m != "" {
		db = db.Where("middle_name_th ILIKE ? OR middle_name_en ILIKE ?", "%"+m+"%", "%"+m+"%")
	}
	if l := filters["last_name"]; l != "" {
		db = db.Where("last_name_th ILIKE ? OR last_name_en ILIKE ?", "%"+l+"%", "%"+l+"%")
	}
	if dob := filters["date_of_birth"]; dob != "" {
		db = db.Where("date_of_birth = ?", dob)
	}
	if phone := filters["phone_number"]; phone != "" {
		db = db.Where("phone_number = ?", phone)
	}
	if email := filters["email"]; email != "" {
		db = db.Where("email = ?", email)
	}

	if err := db.Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}
