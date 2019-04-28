package fs

import "github.com/yidane/formula/opt"

type CosFunction struct {
}

func (*CosFunction) Name() string {
	return "cos"
}

func (*CosFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewCosFunction() *CosFunction {
	return &CosFunction{}
}
