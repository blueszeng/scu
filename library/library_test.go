package library

import (
	"log"
	"testing"

	"github.com/gocolly/colly"
	"github.com/mohuishou/scu/test"
)

func TestNewLibrary(t *testing.T) {
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
			name: "登陆成功",
			args: args{
				studentID: test.LibStudentID,
				password:  test.LibPassword,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLibrary(tt.args.studentID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLibrary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(*got)
		})
	}
}

func TestLibrary_GetLoanAll(t *testing.T) {
	lib, _ := NewLibrary(test.LibStudentID, test.LibPassword)
	type args struct {
		URL string
		c   *colly.Collector
	}
	tests := []struct {
		name string
		lib  *Library
	}{
		{
			name: "获取借阅书籍",
			lib:  lib,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lib.GetLoanAll()
			log.Println(got)
		})
	}
}
