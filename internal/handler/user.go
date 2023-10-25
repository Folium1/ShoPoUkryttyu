package handler

import (
	"log/slog"

	"github.com/Folium1/ShoPoUkryttyu/internal/config"
	"github.com/Folium1/ShoPoUkryttyu/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceInterface interface {
	CreateUser(user models.User) (primitive.ObjectID, error)
	LoginUser(user models.User) (primitive.ObjectID, error)
	GetAllUsers() ([]models.User, error)
}

func SignUpHandler(service UserServiceInterface, cfg config.Config, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const op = "handler.UserRegister"

		var newUser models.User
		if err := c.BodyParser(&newUser); err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusBadRequest).JSON(models.SetResponse(fiber.StatusBadRequest, "Bad request"))
		}

		if err := validateUser(&newUser); err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusBadRequest).JSON(models.SetResponse(fiber.StatusBadRequest, "Bad request"))
		}

		id, err := service.CreateUser(newUser)
		if err != nil {
			log.Error(op, err)
			switch err {
			case models.ErrInvalidCredentials:
				log.Warn(err.Error())
				return c.Status(fiber.StatusBadRequest).JSON(models.SetResponse(fiber.StatusBadRequest, "Bad request"))
			case models.ErrServer:
				log.Error(op, err)
				return c.Status(fiber.StatusInternalServerError).JSON(models.SetResponse(fiber.StatusInternalServerError, "Internal server error"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.SetResponse(fiber.StatusInternalServerError, "Internal server error"))
		}

		token, err := generateToken(id.String(), cfg)
		if err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.SetResponse(fiber.StatusInternalServerError, "Internal server error"))
		}

		return authUser(c, token)
	}
}

func LoginHandler(service UserServiceInterface, cfg config.Config, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const op = "handler.UserLogin"

		var user models.User
		if err := c.BodyParser(&user); err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusBadRequest).JSON(models.SetResponse(fiber.StatusBadRequest, "Bad request"))
		}

		if err := validateUser(&user); err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusBadRequest).JSON(models.SetResponse(fiber.StatusBadRequest, "Bad request"))
		}

		id, err := service.LoginUser(user)
		if err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.SetResponse(fiber.StatusInternalServerError, "Internal server error"))
		}

		token, err := generateToken(id.String(), cfg)
		if err != nil {
			log.Error(op, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.SetResponse(fiber.StatusInternalServerError, "Internal server error"))
		}

		return authUser(c, token)
	}
}

func AllUsers(service UserServiceInterface, cfg config.Config, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const op = "handler.AllUsers"
		users, err := service.GetAllUsers()
		if err != nil {
			log.Error(op, err)
			return c.Status(500).JSON(models.SetResponse(500, "Internal server error"))
		}

		return c.Status(fiber.StatusOK).JSON(users)
	}
}
