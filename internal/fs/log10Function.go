package fs

import "github.com/yidane/formula/opt"

type Log10Function struct {
}

func (*Log10Function) Name() string {
	panic("implement me")
}

func (*Log10Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewLog10Function() *Log10Function {
	return &Log10Function{}
}
