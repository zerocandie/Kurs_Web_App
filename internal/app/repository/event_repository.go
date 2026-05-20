package repository

import (
	"WebApp/internal/app/ds"
	"time"
)



func (r *Repository) GetEventsFiltered(
	title string,
	cityID string,
	dateFrom string,
	dateTo string,
) ([]ds.Event, error) {

	var events []ds.Event

	query := r.db.Model(&ds.Event{}).
		Preload("City").
		Where("is_active = ?", true)

	if title != "" {
		query = query.Where(
			"LOWER(title) LIKE LOWER(?)",
			"%"+title+"%",
		)
	}

	if cityID != "" {
		query = query.Where("city_id = ?", cityID)
	}

	if dateFrom != "" {
		from, err := time.Parse("2006-01-02", dateFrom)
		if err == nil {
			query = query.Where("event_date >= ?", from)
		}
	}

	if dateTo != "" {
		to, err := time.Parse("2006-01-02", dateTo)
		if err == nil {
			query = query.Where("event_date <= ?", to)
		}
	}

	err := query.
		Order("event_date ASC").
		Find(&events).Error

	return events, err
}

func (r *Repository) GetEventByID(id uint) (*ds.Event, error) {
	var event ds.Event

	err := r.db.
		Preload("City").
		First(&event, id).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *Repository) CreateEvent(
	event *ds.Event,
) error {
	return r.db.Create(event).Error
}