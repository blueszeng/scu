package jwc

import (
	"testing"

	"github.com/mohuishou/scu/test"
)

func TestLogin(t *testing.T) {
	type args struct {
		studentID string
		password  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "登录成功！",
			args: args{
				studentID: test.StudentID,
				password:  test.Password,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Login(tt.args.studentID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
