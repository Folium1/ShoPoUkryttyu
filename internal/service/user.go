package service

import (
	"github.com/Folium1/ShoPoUkryttyu/internal/models"
	st "github.com/Folium1/ShoPoUkryttyu/internal/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	storage st.DB
}

func NewUserService(storage *st.DB) *UserService {
	return &UserService{storage: *storage}
}

func (s UserService) CreateUser(user models.User) (primitive.ObjectID, error) {
	var err error
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return primitive.NilObjectID, models.ErrServer
	}

	id, err := s.storage.CreateUser(user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (s UserService) LoginUser(user models.User) (primitive.ObjectID, error) {
	dbUser, err := s.storage.UserByMail(user.Email)
	if err != nil {
		return primitive.NilObjectID, models.ErrInvalidCredentials
	}

	if err := comparePasswords(dbUser.Password, user.Password); err != nil {
		return primitive.NilObjectID, models.ErrInvalidCredentials
	}

	return dbUser.ID, nil
}

func (s UserService) UserByID(id primitive.ObjectID) (models.User, error) {
	user, err := s.storage.UserByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.storage.AllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
