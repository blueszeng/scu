package grade

import (
	"log"
	"testing"

	"github.com/gocolly/colly"
	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"
)

func TestGetNow(t *testing.T) {
	c, _ := jwc.Login(test.StudentID, test.Password)
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grades := GetNow(tt.args.c)
			log.Println(grades)
		})
	}
}

func TestGetALL(t *testing.T) {
	c, _ := jwc.Login(test.StudentID, test.Password)
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetALL(tt.args.c)
			log.Println(got)
		})
	}
}
