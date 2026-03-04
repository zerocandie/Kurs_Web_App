package repository

import (
	"WebApp/internal/app/ds"
)

func (r *Repository) GetAllChats() ([]ds.Event, error) {
	var events []ds.Event
	// Добавили Preload, чтобы подтянулся город
	err := r.db.Where("is_active = true").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *Repository) GetAllChatsById(id int) (*ds.Event, error) {
	var event ds.Event
	// Добавили Preload
	err := r.db.Preload("City").First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *Repository) GetEventsByTitle(title string) ([]ds.Event, error) {
	var events []ds.Event
	err := r.db.Where("title = ?", title).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *Repository) CreateEvent(event ds.Event) (*ds.Event, error) {
	err := r.db.Create(&event).Error
	return &event, err
}
