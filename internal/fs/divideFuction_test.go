package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestDivideFunction_Evaluate(t *testing.T) {
	tests := []struct {
		arg0    string
		arg1    string
		want    float64
		wantErr bool
	}{
		{"3", "2", 1.5, false},
		{"4", "2", 2, false},
		{"5", "3", 5.0 / 3.0, false},
		{"3", "0", 0, true},
	}
	d := NewDivideFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arg0, tt.arg1, tt.want), func(t *testing.T) {
			got, err := d.Evaluate(nil, []*opt.LogicalExpression{exp.NewFloatExpression(tt.arg0), exp.NewFloatExpression(tt.arg1)}...)
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
				t.Fatalf("%v!=%v", tt.want, v)
			}
		})
	}
}
