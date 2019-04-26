package opt

type TernaryExpression struct {
	Left *LogicalExpression
	Middle *LogicalExpression
	Right *LogicalExpression
}

func NewTernaryExpression(left,mid,right *LogicalExpression) *LogicalExpression{
	var result LogicalExpression= &TernaryExpression{
		Left:left,
		Middle:mid,
		Right:right,
	}

	return &result
}

func (*TernaryExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*TernaryExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*TernaryExpression) ToString() string {
	panic("implement me")
}
