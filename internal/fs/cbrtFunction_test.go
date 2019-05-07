package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestCbrtFunction_Evaluate(t *testing.T) {
	tests := []struct {
		arg  string
		want float64
	}{
		{"8", 2},
		{"27", 3},
		{"1", 1},
		{"1.331", 1.1},
	}
	f := NewCbrtFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arg, tt.want), func(t *testing.T) {
			got, err := f.Evaluate(nil, []*opt.LogicalExpression{exp.NewFloatExpression(tt.arg)}...)
			if err != nil {
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
