package fs

import "github.com/yidane/formula/opt"

type CeilingFunction struct {
}

func (*CeilingFunction) Name() string {
	panic("implement me")
}

func (*CeilingFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewCeilingFunction() *CeilingFunction {
	return &CeilingFunction{}
}
