package fs

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type InFunction struct {
}

func (*InFunction) Name() string {
	return "in"
}

func (f *InFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("function %s need more than one arguments", f.Name())
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(args); i++ {
		arg, err := (*args[i]).Evaluate(context)
		if err != nil {
			return nil, err
		}

		if arg0.Equal(arg) {
			return opt.NewArgumentWithType(true, reflect.Bool), nil
		}
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewInFunction() *InFunction {
	return &InFunction{}
}
