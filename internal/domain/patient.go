package domain

import (
	"time"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Patient struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientHN    string     `gorm:"type:varchar(20);not null" json:"patientHn"`
	FirstNameTH  string     `gorm:"type:varchar(100);index:idx_patient_name,priority:1" json:"firstNameTh"`
	MiddleNameTH string     `gorm:"type:varchar(100);index:idx_patient_name,priority:5" json:"middleNameTh"`
	LastNameTH   string     `gorm:"type:varchar(100);index:idx_patient_name,priority:3" json:"lastNameTh"`
	FirstNameEN  string     `gorm:"type:varchar(100);index:idx_patient_name,priority:2" json:"firstNameEn"`
	MiddleNameEN string     `gorm:"type:varchar(100);index:idx_patient_name,priority:6" json:"middleNameEn"`
	LastNameEN   string     `gorm:"type:varchar(100);index:idx_patient_name,priority:4" json:"lastNameEn"`
	DateOfBirth  time.Time  `gorm:"type:date;index:idx_patient_date_of_birth" json:"dateOfBirth"`
	NationalID   string     `gorm:"type:char(13);index:idx_patient_national_id" json:"nationalId"`
	PassportID   string     `gorm:"type:varchar(9);index:idx_patient_passport_id" json:"passportId"`
	PhoneNumber  string     `gorm:"type:varchar(15);index:idx_patient_phone" json:"phoneNumber"`
	Email        string     `gorm:"type:varchar(320);unique;index:idx_patient_email" json:"email"`
	Gender       Gender     `gorm:"type:gender_type" json:"gender"`
	HospitalID   int        `gorm:"not null;index:idx_patient_hospital_id" json:"-"`
	Hospital     *Hospital  `gorm:"foreignKey:HospitalID;constraint:OnDelete:CASCADE" json:"hospital,omitempty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime:false" json:"-"`
}

func (Patient) TableName() string {
	return "patient"
}
