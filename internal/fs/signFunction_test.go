package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"math"
	"reflect"
	"strconv"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestSignFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{[]string{"1", "2", "3", "30", "0.5"}},
	}

	f := NewSignFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				if result.Type != reflect.Bool {
					t.Fatal("error type")
				}

				if v1, _ := strconv.ParseFloat(tt.args[i], 64); math.Signbit(v1) != result.Value.(bool) {
					t.Fatal()
				}
			}
		})
	}
}
