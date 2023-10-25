package handler

import (
	"errors"

	"github.com/Folium1/ShoPoUkryttyu/internal/models"
)

func validateUser(user *models.User) error {
	if len(user.Name) < 3 {
		return errors.New("too short name")
	}
	if user.Email == "" {
		return errors.New("empty email")
	}
	if len(user.Password) < 6 {
		return errors.New("too short password")
	}
	return nil
}
