package fs

import "github.com/yidane/formula/opt"

type RoundFunction struct {
	opt.BaseFunction
}

func (f RoundFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

