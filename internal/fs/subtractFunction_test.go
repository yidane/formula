package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestSubtractFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args    []string
		want    float64
		wantErr bool
	}{
		{[]string{"1", "2"}, -1, false},
		{[]string{"1", "2", "3"}, 3, true},
	}

	p := NewSubtractFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			var args []*opt.LogicalExpression
			for i := 0; i < len(tt.args); i++ {
				args = append(args, exp.NewFloatExpression(tt.args[i]))
			}

			if tt.wantErr && len(tt.args) < 2 {
				args = append(args, exp.NewStringValueExpression("a"))
			}

			got, err := p.Evaluate(nil, args...)
			if err != nil {
				if tt.wantErr {
					t.Log(err)
					return
				}
				t.Fatal(err)
			}

			v, err := got.Float64()
			if err != nil {
				t.Fatal(err)
			}

			if v != tt.want {
				t.Fatalf("%v!+%v", tt.want, v)
			}
		})
	}
}
