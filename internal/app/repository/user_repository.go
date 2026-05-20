package repository

import "WebApp/internal/app/ds"

func (r *Repository) CreateUser(
	user *ds.User,
) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(
	login string,
) (*ds.User, error) {

	var user ds.User

	err := r.db.
		Where(
			"login = ?",
			login,
		).
		First(&user).Error

	return &user, err
}