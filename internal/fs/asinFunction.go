package fs

import "github.com/yidane/formula/opt"

type AsinFunction struct {
	opt.BaseFunction
}

func (f AsinFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

