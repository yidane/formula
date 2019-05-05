package main

import (
	"github.com/yidane/formula/opt"
	"reflect"
)

type CustomFunction struct {
}

func (*CustomFunction) Name() string {
	return "CustomFunction"
}

func (f *CustomFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	left, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	leftValue, err := left.Int64()
	if err != nil {
		return nil, err
	}

	right, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	rightValue, err := right.Int64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(leftValue+rightValue+1, reflect.Int64), nil
}
