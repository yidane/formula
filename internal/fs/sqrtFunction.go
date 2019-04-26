package fs

import "github.com/yidane/formula/opt"

type SqrtFunction struct {
	opt.BaseFunction
}

func (f SqrtFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

