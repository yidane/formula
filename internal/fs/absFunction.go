package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type AbsFunction struct {
}

func (*AbsFunction) Name() string {
	return "Abs"
}

func (f *AbsFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchOneArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Abs(v0), reflect.Float64), nil
}

func NewAbsFunction() *AbsFunction {
	return &AbsFunction{}
}
