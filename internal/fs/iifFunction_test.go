package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestIIFFunction_Evaluate(t *testing.T) {
	type args struct {
		arg0 bool
		arg1 string
		arg2 string
	}
	tests := []args{
		{true, "1", "2"},
		{false, "1", "2"},
	}

	i := NewIIFFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arg0), func(t *testing.T) {
			exp0 := exp.NewBooleanValueExpression(tt.arg0)
			exp1 := exp.NewStringValueExpression(tt.arg1)
			exp2 := exp.NewStringValueExpression(tt.arg2)

			result, err := i.Evaluate(nil, []*opt.LogicalExpression{exp0, exp1, exp2}...)
			if err != nil {
				t.Fatal(err)
			}

			r := tt.arg1
			if !tt.arg0 {
				r = tt.arg2
			}

			if result.String() != r {
				t.Fatal()
			}
		})
	}
}
