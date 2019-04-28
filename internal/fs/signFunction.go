package fs

import "github.com/yidane/formula/opt"

type SignFunction struct {
}

func (*SignFunction) Name() string {
	panic("implement me")
}

func (*SignFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewSignFunction() *SignFunction {
	return &SignFunction{}
}
