package domain

import (
	"time"
)

type Hospital struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string     `gorm:"type:varchar(20);unique;not null" json:"code"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:false" json:"-"`

	Staffs   []Staff   `gorm:"constraint:OnDelete:CASCADE" json:"staffs,omitempty"`
	Patients []Patient `gorm:"constraint:OnDelete:CASCADE" json:"patients,omitempty"`
}

func (Hospital) TableName() string {
	return "hospital"
}
