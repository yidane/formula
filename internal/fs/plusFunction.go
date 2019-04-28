package fs

import "github.com/yidane/formula/opt"

type PlusFunction struct {
}

func (*PlusFunction) Name() string {
	panic("implement me")
}

func (*PlusFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewPlusFunction() *PlusFunction {
	return &PlusFunction{}
}
