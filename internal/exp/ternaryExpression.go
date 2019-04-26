package exp

import "github.com/yidane/formula/opt"

type TernaryExpression struct {
	Left *opt.LogicalExpression
	Middle *opt.LogicalExpression
	Right *opt.LogicalExpression
}

func NewTernaryExpression(left,mid,right *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &TernaryExpression{
		Left:left,
		Middle:mid,
		Right:right,
	}

	return &result
}

func (*TernaryExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*TernaryExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*TernaryExpression) ToString() string {
	panic("implement me")
}
