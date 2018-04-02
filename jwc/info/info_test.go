package info

import (
	"testing"

	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"
)

func TestGet(t *testing.T) {
	c, err := jwc.Login(test.StudentID, test.Password)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Get(c))
}
