package fs

import "github.com/yidane/formula/opt"

type AbsFunction struct {
	opt.BaseFunction
}

func (f AbsFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}



