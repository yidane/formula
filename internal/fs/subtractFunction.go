package fs

import "github.com/yidane/formula/opt"

type SubtractFunction struct {
	opt.BaseFunction
}

func (f SubtractFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}
 
