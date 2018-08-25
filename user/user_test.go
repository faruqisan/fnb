package user

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want *Model
	}{
		{
			name: "Test Success",
			args: args{
				name:  "John Doe",
				email: "john@doe.com",
			},
			want: &Model{
				Name:  "John Doe",
				Email: "john@doe.com",
			},
		},
		{
			name: "Test Fail",
			args: args{
				name:  "",
				email: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.name, tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
