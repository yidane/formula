package fs

import "github.com/yidane/formula/opt"

type AtanFunction struct {
	opt.BaseFunction
}

func (f AtanFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}
