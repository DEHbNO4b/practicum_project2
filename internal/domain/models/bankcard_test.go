package models

import "testing"

func TestCard_SetCardID(t *testing.T) {
	type args struct {
		id []rune
	}
	tests := []struct {
		name    string
		c       *Card
		args    args
		wantErr bool
	}{
		{
			name:    "positive case",
			c:       &Card{},
			args:    args{id: []rune("2992922292234321")},
			wantErr: false,
		},
		{
			name:    "wrong len, negative case",
			c:       &Card{},
			args:    args{id: []rune("299292292234321")},
			wantErr: true,
		},
		{
			name:    "not digit, negative case",
			c:       &Card{},
			args:    args{id: []rune("29929222b2234321")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SetCardID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Card.SetCardID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
