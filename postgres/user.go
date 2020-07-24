package postgres

import (
	"errors"
	"github.com/go-pg/pg/v9"
	"todo/domain"
)

// UserRepo DB pointer
type UserRepo struct {
	DB *pg.DB
}

// GetByEmail get user by email
func (u *UserRepo) GetByEmail(email string) (*domain.User, error) {
	user := new(domain.User)

	err := u.DB.Model(user).Where("email = ?", email).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return user, nil
}

// GetByUsername get user by username
func (u *UserRepo) GetByUsername(username string) (*domain.User, error) {
	user := new(domain.User)

	err := u.DB.Model(user).Where("username = ?", username).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return user, nil
}

// Create new user
func (u *UserRepo) Create(user *domain.User) (*domain.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// NewUserRepo get DB pointer
func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}
