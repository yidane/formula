package opt

type FunctionExpression struct {
	BaseExpression
}

func NewFunctionExpression(id *IdentifierExpression, args []*LogicalExpression) *LogicalExpression{
	var result LogicalExpression=&FunctionExpression{}

	return &result
}

func (*FunctionExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*FunctionExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*FunctionExpression) ToString() string {
	panic("implement me")
}


