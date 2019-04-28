package fs

import "github.com/yidane/formula/opt"

type AsinFunction struct {
}

func (*AsinFunction) Name() string {
	panic("implement me")
}

func (*AsinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewAsinFunction() *AsinFunction{
	return &AsinFunction{}
}

