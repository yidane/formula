package fs

import "github.com/yidane/formula/opt"

type DivideFunction struct {
	opt.BaseFunction
}

func (f DivideFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

