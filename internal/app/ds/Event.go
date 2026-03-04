package ds

import "time"

type Event struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement"`
	Title               string    `gorm:"type:varchar(100);not null"`
	Img                 string    `gorm:"type:varchar(500)"`
	About               string    `gorm:"type:text"`
	Result              string    `gorm:"type:text"`
	EventDate           time.Time `gorm:"not null;index"`
	Location            string    `gorm:"type:varchar(255)"`
	CityID              uint      `gorm:"not null;index"`
	City                City      `gorm:"foreignKey:CityID"`
	OrganizerID         *uint     `gorm:"index"`                  // 🔥 Pointer + index
	Organizer           *Person   `gorm:"foreignKey:OrganizerID"` // 🔥 Pointer
	MaxParticipants     int       `gorm:"default:0"`
	CurrentParticipants int       `gorm:"default:0"`
	IsActive            bool      `gorm:"default:true;index"`
	CreatedAt           time.Time `gorm:"autoCreateTime"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`
}
