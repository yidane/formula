package fs

import (
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestConcatFunction_Evaluate(t *testing.T) {
	want := "123.123"
	c := NewConcatFunction()

	got, err := c.Evaluate(nil, []*opt.LogicalExpression{
		exp.NewIntegerValueExpression("1"),
		exp.NewIntegerValueExpression("2"),
		exp.NewIntegerValueExpression("3"),
		exp.NewStringValueExpression(".123")}...)
	if err != nil {
		t.Fatal(err)
	}

	v := got.String()

	if v != want {
		t.Fatalf("%s!=%s", v, want)
	}

}
