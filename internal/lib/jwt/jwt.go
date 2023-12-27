package jwt

import (
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = user.Login()
	claims["user_id"] = user.ID()
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(app.Secret()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
