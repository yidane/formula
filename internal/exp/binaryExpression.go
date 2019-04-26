package exp

import "github.com/yidane/formula/opt"

type BinaryExpression struct {
	opt.BaseExpression
	Name            string
	LeftExpression  *opt.LogicalExpression
	RightExpression *opt.LogicalExpression
}

func NewBinaryExpression(name string,leftExpression ,rightExpression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression =  &BinaryExpression{
		Name:name,
		LeftExpression:leftExpression,
		RightExpression:rightExpression,
	}

	return &result
}

func (expression *BinaryExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	expression.Context=context

	//p:= unsafe.Pointer(expression)

	return nil
}

func (expression *BinaryExpression) Evaluate() *opt.Argument {
	return opt.NewArgument(nil)
}

func (expression *BinaryExpression) ToString() string {
	return ""
}