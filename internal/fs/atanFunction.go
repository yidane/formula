package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type AtanFunction struct {
}

func (*AtanFunction) Name() string {
	return "atan"
}

func (f *AtanFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Atan(v), reflect.Float64), nil
}

func NewAtanFunction() *AtanFunction {
	return &AtanFunction{}
}
