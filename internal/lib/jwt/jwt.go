package jwt

import (
	"path/filepath"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
	Login  string
}

func BuildJWTString(u models.User) (string, error) {
	cfg := config.MustLoadServCfg()
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.TokenTTL)),
		},
		UserID: u.ID(),
		Login:  u.Login(),
		// собственное утверждение
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetClaims(token string) (Claims, error) {

	confg := config.MustLoadByPath(filepath.FromSlash("./config/server.yaml"))
	// создаём экземпляр структуры с утверждениями
	claims := Claims{}
	// s := strings.Fields(token)

	// if len(s) != 2 {
	// 	return claims, errors.New("wrong authorization field")
	// }
	// парсим из строки токена tokenString в структуру claims
	jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(confg.SecretKey), nil
	})

	// возвращаем ID пользователя в читаемом виде
	return claims, nil
}

// func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			// когда создан токен
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
// 		},
// 		Login:  user.Login(),
// 		UserID: user.ID(),
// 	})

// 	// claims := token.Claims.(jwt.MapClaims)
// 	// claims["login"] = user.Login()
// 	// claims["userID"] = user.ID()
// 	// claims["exp"] = time.Now().Add(duration).Unix()

// 	cfg := config.MustLoadServCfg()

// 	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }
// func GetClaims(token string) (Claims, error) {
// 	// создаём экземпляр структуры с утверждениями
// 	claims := Claims{}
// 	s := strings.Fields(token)

// 	if len(s) != 2 {
// 		return claims, errors.New("wrong authorization field")
// 	}
// 	// парсим из строки токена tokenString в структуру claims
// 	jwt.ParseWithClaims(s[1], &claims, func(t *jwt.Token) (interface{}, error) {
// 		// return []byte(app.Secret()), nil
// 		return []byte("secret_string"), nil
// 	})

// 	// возвращаем ID пользователя в читаемом виде
// 	return claims, nil
// }
