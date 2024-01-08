package tests

import (
	"testing"

	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/DEHbNO4b/practicum_project2/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestSaveBynary_HappyPath(t *testing.T) {

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

	_, err = st.JWTClient.SaveBinary(ctx, &pb.BinaryData{
		Data: []byte("some random bytes"),
		Info: gofakeit.Sentence(10),
	})
	require.NoError(t, err)
	// _, err = st.JWTClient.SaveBinary(ctx, getRandomBinaryData(300))
	// require.NoError(t, err)

	// res, err := st.JWTClient.ShowData(ctx, &pb.Empty{})
	// require.NoError(t, err)

	// fmt.Printf("response log-pass: %+v", res)
}
