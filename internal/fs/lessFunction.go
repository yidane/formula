package fs

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type LessFunction struct {
}

func (*LessFunction) Name() string {
	return "<"
}

func (f *LessFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if !arg0.IsNumber() {
		return nil, fmt.Errorf("function %s required numbers", f.Name())
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, fmt.Errorf("function %s required numbers", f.Name())
	}

	v0, _ := arg0.Float64()
	v1, _ := arg1.Float64()

	if v0 < v1 {
		return opt.NewArgumentWithType(true, reflect.Bool), nil
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewLessFunction() *LessFunction {
	return &LessFunction{}
}
