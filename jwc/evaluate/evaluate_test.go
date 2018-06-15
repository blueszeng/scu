package evaluate

import (
	"testing"
	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"
)

func TestGetEvaList(t *testing.T) {
	c, _ := jwc.Login(test.StudentID, test.Password)
	res,err:=GetEvaList(c)
	t.Log(len(res),err,res)
	//r := res[len(res)-1]
	//r.Comment = "超级棒的老师"
	//r.Star = 5
	//t.Log(AddEvaluate(c,&r))
}
