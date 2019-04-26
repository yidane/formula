package fs

import "github.com/yidane/formula/opt"

type GreaterFunction struct {
	opt.BaseFunction
}

func (f GreaterFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

