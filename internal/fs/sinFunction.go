package fs

import "github.com/yidane/formula/opt"

type SinFunction struct {
	opt.BaseFunction
}

func (f SinFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

