package dto

import "github.com/go-playground/validator/v10"

type PatientFilter struct {
	ID          string `form:"id"`
	FirstName   string `form:"firstName" validate:"omitempty,max=255"`
	MiddleName  string `form:"middleName" validate:"omitempty,max=255"`
	LastName    string `form:"lastName" validate:"omitempty,max=255"`
	DateOfBirth string `form:"dateOfBirth" validate:"omitempty,datetime=2006-01-02"`
	PhoneNumber string `form:"phoneNumber" validate:"omitempty,max=20"`
	Email       string `form:"email" validate:"omitempty,email"`
}

func (r *PatientFilter) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
