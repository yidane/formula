package fs

import "github.com/yidane/formula/opt"

type SignFunction struct {
	opt.BaseFunction
}

func (f SignFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

