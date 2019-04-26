package formula

import (
	"github.com/yidane/formula/opt"
	"testing"
)

func TestNewExpression(t *testing.T) {
	const expString  = "1+2"

	exp:=NewExpression(expString,opt.IgnoreCase)

	result,err:= exp.Evaluate()
	if err!=nil{
		t.Fatal(err)
	}

	if result.(int64)!=3{
		t.Fatal()
	}
}
