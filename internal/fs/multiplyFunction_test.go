package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestMultiplyFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args    []*opt.LogicalExpression
		want    float64
		wantErr bool
	}{
		{[]*opt.LogicalExpression{
			exp.NewFloatExpression("3"),
			exp.NewFloatExpression("2"),
		}, 6, false},
		{[]*opt.LogicalExpression{
			exp.NewFloatExpression("3"),
			exp.NewFloatExpression("2"),
			exp.NewFloatExpression("3"),
			exp.NewFloatExpression("2"),
		}, 0, true},
		{[]*opt.LogicalExpression{
			exp.NewFloatExpression("3"),
			exp.NewStringValueExpression("a"),
		}, 0, true},
	}

	m := NewMultiplyFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			got, err := m.Evaluate(nil, tt.args...)

			if err != nil {
				if tt.wantErr {
					return
				}

				t.Fatal(err)
			}

			v, err := got.Float64()
			if err != nil {
				t.Fatal(err)
			}

			if v != tt.want {
				t.Fatalf("%v!=%v", v, tt.want)
			}
		})
	}
}
