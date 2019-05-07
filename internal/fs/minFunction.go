package fs

import (
	"github.com/yidane/formula/opt"
	"reflect"
)

type MinFunction struct {
}

func (*MinFunction) Name() string {
	return "min"
}

func (f *MinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	result, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	var arg *opt.Argument
	var v float64
	for i := 1; i < len(args); i++ {
		arg, err = (*args[i]).Evaluate(context)
		if err != nil {
			return nil, err
		}
		v, err = arg.Float64()
		if err != nil {
			return nil, err
		}

		if v < result {
			result = v
		}
	}

	return opt.NewArgumentWithType(result, reflect.Float64), nil
}

func NewMinFunction() *MinFunction {
	return &MinFunction{}
}
