package fs

import "github.com/yidane/formula/opt"

type MaxFunction struct {
}

func (*MaxFunction) Name() string {
	panic("implement me")
}

func (*MaxFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewMaxFunction() *MaxFunction {
	return &MaxFunction{}
}
