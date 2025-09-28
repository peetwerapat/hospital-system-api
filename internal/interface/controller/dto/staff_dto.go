package dto

import "github.com/go-playground/validator/v10"

type StaffRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	HospitalID int    `json:"hospitalId" binding:"required"`
}

func (r *StaffRequest) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
