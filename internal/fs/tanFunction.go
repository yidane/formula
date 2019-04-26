package fs

import "github.com/yidane/formula/opt"

type TanFunction struct {
	opt.BaseFunction
}

func (f TanFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

