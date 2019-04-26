package fs

import "github.com/yidane/formula/opt"

type TruncateFunction struct {
	opt.BaseFunction
}

func (f TruncateFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

