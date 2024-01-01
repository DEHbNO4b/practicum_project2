package tests

import (
	"testing"

	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/DEHbNO4b/practicum_project2/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShowData_HappyPath(t *testing.T) {

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

	err = st.MakeJWTClient(token)
	require.NoError(t, err)

	_, err = st.JWTClient.SaveText(ctx, getRandomTextData())
	assert.NoError(t, err)
	_, err = st.JWTClient.SaveText(ctx, getRandomTextData())
	assert.NoError(t, err)

	_, err = st.JWTClient.SaveLogPass(ctx, getRandomLogPassData())
	assert.NoError(t, err)
	_, err = st.JWTClient.SaveLogPass(ctx, getRandomLogPassData())
	assert.NoError(t, err)
	_, err = st.JWTClient.SaveLogPass(ctx, getRandomLogPassData())
	assert.NoError(t, err)

	res, err := st.JWTClient.ShowData(ctx, &pb.Empty{})

	assert.NoError(t, err)
	assert.NotEmpty(t, res.GetLpd())
	assert.NotEmpty(t, res.GetTd())
	assert.Equal(t, 2, len(res.GetTd()))
	assert.Equal(t, 3, len(res.GetLpd()))

}
