package ecard

import (
	"testing"
	"time"

	"github.com/mohuishou/scu"
	"github.com/mohuishou/scu/test"

	"github.com/gocolly/colly"
)

func TestGet(t *testing.T) {
	c, _ := scu.NewCollector(test.StudentID, test.Password)
	start, _ := time.Parse("2006-01-02", "2017-10-10")
	type args struct {
		c     *colly.Collector
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "获取一卡通数据",
			args: args{
				c:     c,
				start: start,
				end:   time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Get(tt.args.c, tt.args.start, tt.args.end))
		})
	}
}
