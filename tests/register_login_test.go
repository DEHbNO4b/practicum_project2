package tests

import (
	"testing"

	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
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

	respReg, err := st.Client.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	require.NoError(t, err)

	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.Client.Login(ctx, &pb.AuthInfo{
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
	// assert.Equal(t, re)
}
func TestDoubleRegister(t *testing.T) {
	ctx, st := suite.New(t)

	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	_, err := st.Client.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})
	require.NoError(t, err)

	_, err = st.Client.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	require.Error(t, err)

	assert.ErrorContains(t, err, "user already exists")

}

func TestLoginError(t *testing.T) {
	ctx, st := suite.New(t)

	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	respReg, err := st.Client.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	require.NoError(t, err)

	assert.NotEmpty(t, respReg.GetUserId())

	_, err = st.Client.Login(ctx, &pb.AuthInfo{
		Login:    login,
		Password: "wrong_pass",
	})

	assert.Error(t, err)

	assert.ErrorContains(t, err, "invalid credentials")
}
