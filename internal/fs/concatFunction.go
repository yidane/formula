package fs

import "github.com/yidane/formula/opt"

type ConcatFunction struct {
	opt.BaseFunction
}

func (f ConcatFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}
