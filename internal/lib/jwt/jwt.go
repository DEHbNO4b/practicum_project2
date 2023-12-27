package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
	Login  string
}

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		Login:  user.Login(),
		UserID: user.ID(),
	})

	// claims := token.Claims.(jwt.MapClaims)
	// claims["login"] = user.Login()
	// claims["userID"] = user.ID()
	// claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(app.Secret()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GetClaims(token string, app models.App) (Claims, error) {
	// создаём экземпляр структуры с утверждениями
	claims := Claims{}
	s := strings.Fields(token)

	if len(s) != 2 {
		return claims, errors.New("wrong authorization field")
	}
	// парсим из строки токена tokenString в структуру claims
	jwt.ParseWithClaims(s[1], &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(app.Secret()), nil
	})

	// возвращаем ID пользователя в читаемом виде
	return claims, nil
}
