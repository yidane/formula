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
	err := opt.MatchOneArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Atan(v), reflect.Float64), nil
}

func NewAtanFunction() *AtanFunction {
	return &AtanFunction{}
}
