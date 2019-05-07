package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestShiftRightFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args    []*opt.LogicalExpression
		want    int64
		wantErr bool
	}{
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("1"),
			exp.NewIntegerValueExpression("0"),
		}, 1, false},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("1"),
			exp.NewIntegerValueExpression("3"),
		}, 1 >> 3, false},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("1"),
			exp.NewIntegerValueExpression("0"),
			exp.NewIntegerValueExpression("0"),
		}, 1, true},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("-1"),
			exp.NewIntegerValueExpression("0"),
		}, -1, false},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("-1"),
			exp.NewIntegerValueExpression("3"),
		}, -1 >> 3, false},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("-1"),
			exp.NewIntegerValueExpression("-1"),
		}, -1, true},
		{[]*opt.LogicalExpression{
			exp.NewFloatExpression("1.1"),
			exp.NewIntegerValueExpression("-1"),
		}, -1, true},
		{[]*opt.LogicalExpression{
			exp.NewIntegerValueExpression("-1"),
			exp.NewFloatExpression("1.1"),
		}, -1, true},
	}

	s := NewShiftRightFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			got, err := s.Evaluate(nil, tt.args...)

			if err != nil {
				if tt.wantErr {
					t.Log(err)
					return
				}

				t.Fatal(err)
			}

			v, err := got.Int64()
			if err != nil {
				t.Fatal(err)
			}

			if v != tt.want {
				t.Fatalf("%v!=%v", v, tt.want)
			}
		})
	}
}
