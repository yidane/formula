package exp

import "github.com/yidane/formula/opt"

type FunctionExpression struct {
	opt.BaseExpression
}

func NewFunctionExpression(id *IdentifierExpression, args []*opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression =&FunctionExpression{}

	return &result
}

func (*FunctionExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*FunctionExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*FunctionExpression) ToString() string {
	panic("implement me")
}


