package tests

import (
	"testing"

	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"github.com/DEHbNO4b/practicum_project2/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	appSecret = "secret_string"
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)
	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	_, err := st.Client.Register(ctx, &pb.RegisterRequest{
		Login:    login,
		Password: pass,
	})
	require.NoError(t, err)

	respLogin, err := st.Client.Login(ctx, &pb.LoginRequest{
		Login:    login,
		Password: pass,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	assert.Equal(t, login, claims["login"].(string))
}
