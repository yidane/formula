package fs

import "github.com/yidane/formula/opt"

type ExpFunction struct {
	opt.BaseFunction
}

func (f ExpFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

