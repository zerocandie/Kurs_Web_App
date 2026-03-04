package ds

import "time"

// Person представляет таблицу persons в базе данных
type Person struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex"`
	Phone     string    `gorm:"type:varchar(20)"`
	CityID    uint      `gorm:"not null;index"`    // 🔥 Внешний ключ
	City      City      `gorm:"foreignKey:CityID"` // 🔥 Связь с таблицей городов
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
