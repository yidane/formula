package fs

import "github.com/yidane/formula/opt"

type MinFunction struct {
	opt.BaseFunction
}

func (f MinFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

