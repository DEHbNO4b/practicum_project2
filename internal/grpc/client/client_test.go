package client

import (
	"context"
	"testing"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/mocks"
	pbkeeper "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/golang/mock/gomock"
)

func TestGophClient_SaveLogPass(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockGophKeeperClient(ctrl)

	// m.EXPECT.SaveLogPass(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)
	m.EXPECT().SaveLogPass(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)

	client := GophClient{
		Ctx:       context.Background(),
		JWTClient: m,
	}

	lp := &models.LogPassData{}
	lp.SetLogin("login")
	lp.SetPass("pass")

	nlp := &models.LogPassData{}

	type args struct {
		ctx context.Context
		lp  *models.LogPassData
	}
	tests := []struct {
		name    string
		g       *GophClient
		args    args
		wantErr bool
	}{
		{
			name: "pozitive case",
			g:    &client,
			args: args{
				ctx: context.Background(),
				lp:  lp,
			},
			wantErr: false,
		},
		{
			name: "negative case",
			g:    &client,
			args: args{
				ctx: context.Background(),
				lp:  nlp,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.SaveLogPass(tt.args.ctx, tt.args.lp); (err != nil) != tt.wantErr {
				t.Errorf("GophClient.SaveLogPass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func client(t *testing.T) *GophClient {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockGophKeeperClient(ctrl)

	// m.EXPECT.SaveLogPass(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)
	m.EXPECT().SaveLogPass(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)
	m.EXPECT().SaveCard(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)

	client := GophClient{
		Ctx:       context.Background(),
		JWTClient: m,
	}
	return &client
}
func TestGophClient_SaveCard(t *testing.T) {

	client := client(t)

	type args struct {
		ctx context.Context
		c   *models.Card
	}

	card := models.Card{}
	card.SetCardID([]rune("1234 1234 1234 1234"))
	card.SetPass("123")
	card.SetDate("2014/03")

	tests := []struct {
		name    string
		g       *GophClient
		args    args
		wantErr bool
	}{
		{
			name: "pozitive case",
			g:    client,
			args: args{
				ctx: context.Background(),
				c:   &card,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.SaveCard(tt.args.ctx, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GophClient.SaveCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
