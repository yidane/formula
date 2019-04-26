package fs

import "github.com/yidane/formula/opt"

type LogFunction struct {
	opt.BaseFunction
}

func (f LogFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

