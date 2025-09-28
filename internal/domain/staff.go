package domain

import (
	"time"
)

type Staff struct {
	ID         int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string     `gorm:"type:varchar(50);not null" json:"username"`
	Password   string     `gorm:"type:varchar(255);not null" json:"-"`
	HospitalID int        `gorm:"not null" json:"-"`
	Hospital   *Hospital  `gorm:"foreignKey:HospitalID;constraint:OnDelete:CASCADE" json:"hospital,omitempty"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime:false" json:"-"`
}

func (Staff) TableName() string {
	return "staff"
}
