package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type AtanFunction struct {
}

func (*AtanFunction) Name() string {
	return "atan"
}

func (*AtanFunction) AliasName() []string {
	return nil
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
