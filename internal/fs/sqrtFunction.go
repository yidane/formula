package fs

import "github.com/yidane/formula/opt"

type SqrtFunction struct {
}

func (*SqrtFunction) Name() string {
	panic("implement me")
}

func (*SqrtFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewSqrtFunction() *SqrtFunction {
	return &SqrtFunction{}
}
