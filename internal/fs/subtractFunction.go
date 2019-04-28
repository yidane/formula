package fs

import "github.com/yidane/formula/opt"

type SubtractFunction struct {
}

func (*SubtractFunction) Name() string {
	panic("implement me")
}

func (*SubtractFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewSubtractFunction() *SubtractFunction {
	return &SubtractFunction{}
}
