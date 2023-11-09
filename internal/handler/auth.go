package handler

import (
	"fmt"
	"time"

	"github.com/Folium1/ShoPoUkryttyu/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func generateToken(userID string, cfg config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	tokenString, err := token.SignedString([]byte(cfg.Salt))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authUser(c *fiber.Ctx, token string) error {
	cookie := &fiber.Cookie{
		Name:     "access",
		Value:    fmt.Sprintf("Bearer %s", token),
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add((24 * time.Hour) * 30),
	}

	c.Cookie(cookie)

	return c.Redirect("/", 302)
}
