package fs

import "github.com/yidane/formula/opt"

const add  = "+"

type AddFunction struct {
	opt.BaseFunction
}

func NewAddFunction() AddFunction{
	var result= AddFunction{}

	result.Name= func() string {
		return add
	}
	result.LowerName= func() string {
		return add
	}

	return result
}

func (f AddFunction) Evaluate(args ...opt.LogicalExpression) *opt.Argument {
	f.MatchArgument(args...)

	return nil
}

