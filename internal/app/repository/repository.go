package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

type Events_City struct {
	ID    int
	Title string
	City  string
}

func (r *Repository) GetEvents() ([]Events_City, error) {
	orders := []Events_City{
		{
			ID:    1,
			Title: "Благоустройство территории ",
			City:  "Оренбург",
		},
		{
			ID:    2,
			Title: "Выборы мэра города Норильск",
			City:  "Норильск",
		},

		{
			ID:    3,
			Title: "Выбор Молодёжной столицы России",
			City:  "Любой",
		},
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("Ошибка, ничего не найдено")
	}
	return orders, nil
}

func (r *Repository) GetCity(id int) (Events_City, error) {
	events, err := r.GetEvents()
	if err != nil {
		return Events_City{}, err
	}

	for _, events := range events {
		if events.ID == id {
			return events, nil
		}
	}
	return Events_City{}, fmt.Errorf("Города не найдены")
}

func (r *Repository) GetCityByTitle(title string) ([]Events_City, error) {
	events, err := r.GetEvents()
	if err != nil {
		return []Events_City{}, err
	}

	var result []Events_City
	for _, order := range events {
		if strings.Contains(strings.ToLower(order.Title), strings.ToLower(title)) {
			result = append(result, order)
		}
	}
	return result, nil
}
