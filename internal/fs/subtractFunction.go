package fs

import (
	"github.com/yidane/formula/opt"
	"reflect"
)

type SubtractFunction struct {
}

func (*SubtractFunction) Name() string {
	return "-"
}

func (*SubtractFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument("+", args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}
	v1, err := arg1.Float64()
	if err != nil {
		return nil, err
	}
	return opt.NewArgumentWithType(v0-v1, reflect.Float64), nil
}

func NewSubtractFunction() *SubtractFunction {
	return &SubtractFunction{}
}
