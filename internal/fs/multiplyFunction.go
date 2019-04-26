package fs

import "github.com/yidane/formula/opt"

type MultiplyFunction struct {
	opt.BaseFunction
}

func (f MultiplyFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

