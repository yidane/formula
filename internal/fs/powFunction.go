package fs

import "github.com/yidane/formula/opt"

type PowerFunction struct {
	opt.BaseFunction
}

func (f PowerFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

