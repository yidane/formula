package fs

import "github.com/yidane/formula/opt"

type PlusFunction struct {
	opt.BaseFunction
}

func (f PlusFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

