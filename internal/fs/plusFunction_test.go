package fs

import (
	"reflect"
	"testing"

	"github.com/yidane/formula/opt"
)

func TestPlusFunction_Evaluate(t *testing.T) {
	type args struct {
		context *opt.FormulaContext
		args    []*opt.LogicalExpression
	}
	tests := []struct {
		name    string
		p       *PlusFunction
		args    args
		want    *opt.Argument
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlusFunction{}
			got, err := p.Evaluate(tt.args.context, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlusFunction.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusFunction.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
