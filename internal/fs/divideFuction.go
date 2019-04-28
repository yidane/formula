package fs

import "github.com/yidane/formula/opt"

type DivideFunction struct {
}

func (*DivideFunction) Name() string {
	panic("implement me")
}

func (*DivideFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewDivideFunction() *DivideFunction {
	return &DivideFunction{}
}
