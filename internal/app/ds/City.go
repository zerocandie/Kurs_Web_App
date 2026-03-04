package ds

import "time"

// City представляет таблицу cities в базе данных
type City struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(50);not null;index"`
	Img       string    `gorm:"type:varchar(255);not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Persons   []Person  `gorm:"foreignKey:CityID"`
	Events    []Event   `gorm:"foreignKey:CityID"`
}
