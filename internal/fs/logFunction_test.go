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
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			f := NewLog10Function()

			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[0])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				if v1, _ := strconv.ParseFloat(tt.args[0], 10); math.Log(v1) != v {
					t.Fatal()
				}
			}
		})
	}
}
