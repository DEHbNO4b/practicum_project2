package jwt

import (
	"testing"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

func TestNewToken(t *testing.T) {

	passHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	u := models.User{}
	u.SetID(1)
	u.SetLogin("first")
	u.SetPassHash(string(passHash))

	app := models.App{}
	app.SetId(1)
	app.SetName("auth")
	app.SetSecret("secret-key")

	type args struct {
		user     models.User
		app      models.App
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				user:     u,
				app:      app,
				duration: 1 * time.Hour,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewToken(tt.args.user, tt.args.app, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("NewToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
