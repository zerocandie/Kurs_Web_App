package ds

import "time"

type Application struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	PersonID    uint      `gorm:"not null;index"`
	Person      Person    `gorm:"foreignKey:PersonID"`
	CityID      uint      `gorm:"not null;index"`
	City        City      `gorm:"foreignKey:CityID"`
	Title       string    `gorm:"type:varchar(150);not null"`
	Description string    `gorm:"type:text;not null"`
	Img         string    `gorm:"type:varchar(255)"`
	Status      string    `gorm:"type:varchar(20);default:'new';index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

const (
	ApplicationStatusNew      = "new"
	ApplicationStatusReview   = "review"
	ApplicationStatusApproved = "approved"
	ApplicationStatusRejected = "rejected"
)
