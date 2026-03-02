package repository

import (
	"WebApp/internal/app/ds"
	"fmt"
)

func (r *Repository) GetCitys() ([]ds.City, error) {
	var citys []ds.City
	err := r.db.Find(&citys).Error
	// обязательно проверяем ошибки, и если они появились - передаем выше, то есть хендлеру
	if err != nil {
		return nil, err
	}
	if len(citys) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return citys, nil
}

func (r *Repository) GetCity(id int) (ds.City, error) {
	city := ds.City{}
	err := r.db.Where("id = ?", id).First(&city).Error
	if err != nil {
		return ds.City{}, err
	}
	return city, nil
}

func (r *Repository) GetCitysByTitle(title string) ([]ds.City, error) {
	var orders []ds.City
	err := r.db.Where("name ILIKE ?", "%"+title+"%").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
