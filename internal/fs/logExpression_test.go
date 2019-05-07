package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"math"
	"strconv"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestLogFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{[]string{"1", "2", "3", "30", "0.5"}},
	}

	f := NewLogFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			var logicalExpression1 = *exp.NewFloatExpression("3")
			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression, &logicalExpression1}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				v1, _ := strconv.ParseFloat(tt.args[i], 64)
				if math.Log(v1)/math.Log(3) != v {
					t.Fatal()
				}
			}
		})
	}
}
