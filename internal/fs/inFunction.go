package fs

import "github.com/yidane/formula/opt"

type InFunction struct {
	opt.BaseFunction
}

func (f InFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

