package fs

import "github.com/yidane/formula/opt"

type InFunction struct {
}

func (*InFunction) Name() string {
	return "in"
}

func (*InFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewInFunction() *InFunction {
	return &InFunction{}
}
