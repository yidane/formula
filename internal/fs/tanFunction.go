package fs

import "github.com/yidane/formula/opt"

type TanFunction struct {
}

func (*TanFunction) Name() string {
	panic("implement me")
}

func (*TanFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewTanFunction() *TanFunction {
	return &TanFunction{}
}
