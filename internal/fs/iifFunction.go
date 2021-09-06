package fs

import (
	"fmt"
	"reflect"

	"github.com/yidane/formula/opt"
)

type IIFFunction struct {
}

func (*IIFFunction) Name() string {
	return "iif"
}

func (*IIFFunction) AliasName() []string {
	return nil
}

func (f *IIFFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("function %s required three arguments", f.Name())
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.Type != reflect.Bool {
		return nil, fmt.Errorf("the first argument of function %s should be bool", f.Name())
	}

	if arg0.Value.(bool) {
		return (*args[1]).Evaluate(context)
	}

	return (*args[2]).Evaluate(context)
}

func NewIIFFunction() *IIFFunction {
	return &IIFFunction{}
}
