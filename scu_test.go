package scu

import (
	"reflect"
	"testing"

	"github.com/gocolly/colly"
)

func TestNewCollector(t *testing.T) {
	type args struct {
		studentID string
		password  string
	}
	tests := []struct {
		name    string
		args    args
		want    *colly.Collector
		wantErr bool
	}{
		{
			name: "登录失败",
			args: args{
				studentID: "34567890",
				password:  "lalalalalal",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCollector(tt.args.studentID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollector() = %v, want %v", got, tt.want)
			}
		})
	}
}
