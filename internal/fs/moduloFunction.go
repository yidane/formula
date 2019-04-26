package fs

import "github.com/yidane/formula/opt"

type ModuloFunction struct {
	opt.BaseFunction
}

func (f ModuloFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	panic("implement me")
}

