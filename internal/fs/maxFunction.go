package fs

import "github.com/yidane/formula/opt"

type MaxFunction struct {
	opt.BaseFunction
}

func (f MaxFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

