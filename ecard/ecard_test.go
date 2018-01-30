package ecard

import (
	"testing"

	"github.com/mohuishou/scu"
	"github.com/mohuishou/scu/test"

	"github.com/gocolly/colly"
)

func TestGet(t *testing.T) {
	c, _ := scu.NewCollector(test.StudentID, test.Password)
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "获取一卡通数据",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Get(tt.args.c))
		})
	}
}
