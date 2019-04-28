package fs

import "github.com/yidane/formula/opt"

type TruncateFunction struct {
}

func (*TruncateFunction) Name() string {
	return "truncate"
}

func (*TruncateFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewTruncateFunction() *TruncateFunction {
	return &TruncateFunction{}
}
