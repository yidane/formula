package opt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgument_Equal(t *testing.T) {
	tests := []struct {
		v    interface{}
		t    reflect.Kind
		want bool
	}{
		{1, reflect.Int, true},
		{1.00, reflect.Float64, true},
		{"1", reflect.String, false},
		{"1.00", reflect.String, false},
	}

	arg := &Argument{
		Value: 1,
		Type:  reflect.Int,
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.v), func(t *testing.T) {
			other := &Argument{
				Value: tt.v,
				Type:  tt.t,
			}
			if got := arg.Equal(other); got != tt.want {
				t.Errorf("Argument.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
