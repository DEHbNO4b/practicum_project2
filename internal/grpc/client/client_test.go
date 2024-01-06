package client

import (
	"context"
	"testing"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	// "github.com/DEHbNO4b/practicum_project2/mocks"

	pbkeeper "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/golang/mock/gomock"
)

func TestGophClient_SaveLogPass(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockGophKeeperClient(ctrl)

	m.EXPECT.SaveLogPass(gomock.Any(), gomock.Any()).Return(&pbkeeper.Empty{}, nil)

	client := GophClient{
		Ctx:       context.Background(),
		JWTClient: m,
	}

	lp := &models.LogPassData{}
	lp.SetLogin("login")
	lp.SetPass("pass")

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
