package fs

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/yidane/formula/internal/exp"

	"github.com/yidane/formula/opt"
)

func TestAcosFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{[]string{"1", "2", "3", "30", "0.5"}},
	}
	f := NewAcosFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				v1, _ := strconv.ParseFloat(tt.args[i], 64)
				r := math.Acos(v1)

				if math.IsNaN(v) && math.IsNaN(r) {
					continue
				}

				if r != v {
					t.Fatal()
				}
			}
		})
	}
}
