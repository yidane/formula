package exp

import "github.com/yidane/formula/opt"

type FunctionExpression struct {
	opt.BaseExpression
}

func NewFunctionExpression(id *IdentifierExpression, args []*opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FunctionExpression{}

	return &result
}

func (*FunctionExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	panic("implement me")
}
