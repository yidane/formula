package fs

import "github.com/yidane/formula/opt"

type IIFFunction struct {
	opt.BaseFunction
}

func (f IIFFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

