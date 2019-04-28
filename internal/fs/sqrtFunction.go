package fs

import "github.com/yidane/formula/opt"

type SqrtFunction struct {
}

func (*SqrtFunction) Name() string {
	return "sqrt"
}

func (*SqrtFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewSqrtFunction() *SqrtFunction {
	return &SqrtFunction{}
}
