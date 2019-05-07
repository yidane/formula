package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestMaxFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args    []string
		want    float64
		wantErr bool
	}{
		{[]string{"1", "2", "3"}, 3, false},
		{[]string{"1", "2"}, 3, true},
	}
	m := NewMaxFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			var args []*opt.LogicalExpression
			for i := 0; i < len(tt.args); i++ {
				args = append(args, exp.NewFloatExpression(tt.args[i]))
			}

			if tt.wantErr {
				args = append(args, exp.NewStringValueExpression("a")) //parse error
			}

			got, err := m.Evaluate(nil, args...)
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
