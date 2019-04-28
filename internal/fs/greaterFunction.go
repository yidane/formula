package fs

import "github.com/yidane/formula/opt"

type GreaterFunction struct {
}

func (*GreaterFunction) Name() string {
	return ">"
}

func (*GreaterFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewGreaterFunction() *GreaterFunction {
	return &GreaterFunction{}
}
