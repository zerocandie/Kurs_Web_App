package repository

import "WebApp/internal/app/ds"

func (r *Repository) GetAllCities() (
	[]ds.City,
	error,
) {

	var cities []ds.City

	err := r.db.Find(&cities).Error

	return cities, err
}