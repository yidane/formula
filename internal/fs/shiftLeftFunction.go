package fs

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type ShiftLeftFunction struct {
}

func (*ShiftLeftFunction) Name() string {
	return "<<"
}

func (f *ShiftLeftFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v0, err := arg0.Int64()
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v1, err := arg1.Int()
	if err != nil {
		return nil, err
	}

	if v1 < 0 {
		return nil, fmt.Errorf("shift bit should greater zero")
	}

	return opt.NewArgumentWithType(v0<<uint(v1), reflect.Int64), nil
}

func NewShiftLeftFunction() *ShiftLeftFunction {
	return &ShiftLeftFunction{}
}
