package middleware

import (
	"strings"

	"github.com/Folium1/ShoPoUkryttyu/internal/config"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
)

const (
	cookieName = "access"
)

func UserIsAuth(cfg config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.Split(ctx.Cookies(cookieName), " ")[1]
		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "empty cookie")
		}

		userID, err := parseToken(token, cfg.Salt)
		if err != nil {
			ctx.ClearCookie(cookieName)
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		ctx.Locals("userID", userID)

		return ctx.Next()
	}
}

func parseToken(tokenString, salt string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	return userId, nil
}
