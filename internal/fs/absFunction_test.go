package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestAbsFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
		want []float64
	}{
		{[]string{"1", "2", "1.1", "-1", "2", "-3"}, []float64{1, 2, 1.1, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			f := &AbsFunction{}

			for i := 0; i < len(tt.args); i++ {
				var logicExpression = *exp.NewIntegerValueExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				if v != tt.want[i] {

				}
			}

		})
	}
}
