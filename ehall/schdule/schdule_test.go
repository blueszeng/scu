package schdule

import (
	"testing"

	"github.com/mohuishou/scu/ehall"
	"github.com/mohuishou/scu/test"
)

func TestGet(t *testing.T) {
	c, err := ehall.Login(test.EhallStudentID, test.EhallPassword)
	if err != nil {
		panic(err)
	}
	t.Log(Get(c, "20172"))
}
