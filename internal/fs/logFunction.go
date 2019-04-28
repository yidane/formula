package fs

import "github.com/yidane/formula/opt"

type LogFunction struct {
}

func (*LogFunction) Name() string {
	panic("implement me")
}

func (*LogFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewLogFunction() *LogFunction {
	return &LogFunction{}
}
