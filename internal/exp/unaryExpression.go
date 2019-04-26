package exp

import "github.com/yidane/formula/opt"

type UnaryExpression struct {
	Name       string
	Expression opt.LogicalExpression
}

func (*UnaryExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*UnaryExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*UnaryExpression) ToString() string {
	panic("implement me")
}

type NotUnaryExpression struct {
	UnaryExpression
}

func NewNotUnaryExpression(name string,expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &NotUnaryExpression{

	}

	return &result
}


func (expression *NotUnaryExpression) Evaluate() *opt.Argument{
	return expression.UnaryExpression.Evaluate()
}

type BitwiseNotUnaryExpression struct {
	UnaryExpression
}

func NewBitwiseNotUnaryExpression(name string,expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BitwiseNotUnaryExpression{
	}

	return &result
}

func (expression *BitwiseNotUnaryExpression) Evaluate() *opt.Argument{
	return expression.UnaryExpression.Evaluate()
}

type NegateUnaryExpression struct {
	UnaryExpression
}

func NewNegateUnaryExpression(name string,expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &NegateUnaryExpression{}

	return &result
}

func (expression *NegateUnaryExpression) Evaluate() *opt.Argument{
	return expression.UnaryExpression.Evaluate()
}