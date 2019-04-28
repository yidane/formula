package fs

import "github.com/yidane/formula/opt"

type ConcatFunction struct {
}

func (*ConcatFunction) Name() string {
	return "concat"
}

func (*ConcatFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewConcatFunction() *ConcatFunction {
	return &ConcatFunction{}
}
