package fs

import (
	"fmt"
	"github.com/yidane/formula/opt"
)

type InFunction struct {
}

func (*InFunction) Name() string {
	return "in"
}

func (f *InFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("function %s need more than tow arguments", f.Name())
	}

	return nil, nil
}

func NewInFunction() *InFunction {
	return &InFunction{}
}
