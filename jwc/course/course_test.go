package course

import (
	"log"
	"net/url"
	"testing"

	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"

	"github.com/gocolly/colly"
)

func TestGet(t *testing.T) {
	c, _ := jwc.Login(test.StudentID, test.Password)
	type args struct {
		c      *colly.Collector
		params url.Values
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "本学期课程测试",
			args: args{
				c:      c,
				params: url.Values{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Get(tt.args.c, tt.args.params)
			log.Println(got)
		})
	}
}
