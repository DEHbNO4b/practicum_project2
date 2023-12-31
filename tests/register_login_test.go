package tests

import (
	"fmt"
	"testing"

	myjwt "github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/DEHbNO4b/practicum_project2/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {

	ctx, st := suite.New(t)

	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	respReg, err := st.AuthClient.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	require.NoError(t, err)

	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	claims, err := myjwt.GetClaims(token)
	require.NoError(t, err)

	assert.Equal(t, login, claims.Login)

}
func TestDoubleRegister(t *testing.T) {
	ctx, st := suite.New(t)

	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	_, err := st.AuthClient.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})
	require.NoError(t, err)

	_, err = st.AuthClient.Register(ctx, &pb.AuthInfo{
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

	respReg, err := st.AuthClient.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	require.NoError(t, err)

	assert.NotEmpty(t, respReg.GetUserId())

	_, err = st.AuthClient.Login(ctx, &pb.AuthInfo{
		Login:    login,
		Password: "wrong_pass",
	})

	assert.Error(t, err)

	assert.ErrorContains(t, err, "invalid credentials")
}

func TestSaveLogPass_HappyPath(t *testing.T) {

	ctx, st := suite.New(t)

	login := gofakeit.Name()
	pass := gofakeit.Password(true, true, true, true, false, 10)

	_, err := st.AuthClient.Register(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})
	if err != nil {
		t.Fatalf("unable to  register  %v", err)
	}

	respLogin, err := st.AuthClient.Login(ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})
	if err != nil {
		t.Fatalf("unable to  login %v", err)
	}

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	fmt.Println("token: ", token)

	err = st.MakeJWTClient(token)
	require.NoError(t, err)

	_, err = st.JWTClient.SaveLogPass(ctx, &pb.SaveLogPassRequest{
		Login:    "saved_login",
		Password: "saved_password",
	})

	require.NoError(t, err)
}
