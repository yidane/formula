package fs

import "github.com/yidane/formula/opt"

type CeilingFunction struct {
}

func (*CeilingFunction) Name() string {
	return "ceiling"
}

func (*CeilingFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewCeilingFunction() *CeilingFunction {
	return &CeilingFunction{}
}
