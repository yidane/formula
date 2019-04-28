package fs

import "github.com/yidane/formula/opt"

type SinFunction struct {
}

func (*SinFunction) Name() string {
	panic("implement me")
}

func (*SinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewSinFunction() *SinFunction{
	return &SinFunction{}
}