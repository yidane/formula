package fs

import "github.com/yidane/formula/opt"

type Log2Function struct {
}

func (*Log2Function) Name() string {
	return "log2"
}

func (*Log2Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewLog2Function() *Log2Function {
	return &Log2Function{}
}
