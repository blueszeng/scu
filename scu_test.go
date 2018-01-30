package scu

import (
	"testing"
)

func TestNewCollector(t *testing.T) {
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
			name: "登录失败",
			args: args{
				studentID: "34567890",
				password:  "lalalalalal",
			},
			wantErr: true,
		},
		{
			name: "登录成功",
			args: args{
				studentID: studentID,
				password:  password,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCollector(tt.args.studentID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
