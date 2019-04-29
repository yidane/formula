package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type TanFunction struct {
}

func (*TanFunction) Name() string {
	return "tan"
}

func (f *TanFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Tan(v), reflect.Float64), nil
}

func NewTanFunction() *TanFunction {
	return &TanFunction{}
}
