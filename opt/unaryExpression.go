package opt

type UnaryExpression struct {
	Name string
	Expression LogicalExpression
}

func (*UnaryExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*UnaryExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*UnaryExpression) ToString() string {
	panic("implement me")
}

type NotUnaryExpression struct {
	UnaryExpression
}

func NewNotUnaryExpression(name string,expression *LogicalExpression) *LogicalExpression{
	var result LogicalExpression= &NotUnaryExpression{

	}

	return &result
}


func (expression *NotUnaryExpression) Evaluate() *Argument{
	return expression.UnaryExpression.Evaluate()
}

type BitwiseNotUnaryExpression struct {
	UnaryExpression
}

func NewBitwiseNotUnaryExpression(name string,expression *LogicalExpression) *LogicalExpression{
	var result LogicalExpression= &BitwiseNotUnaryExpression{
	}

	return &result
}

func (expression *BitwiseNotUnaryExpression) Evaluate() *Argument{
	return expression.UnaryExpression.Evaluate()
}

type NegateUnaryExpression struct {
	UnaryExpression
}

func NewNegateUnaryExpression(name string,expression *LogicalExpression) *LogicalExpression{
	var result LogicalExpression= &NegateUnaryExpression{}

	return &result
}

func (expression *NegateUnaryExpression) Evaluate() *Argument{
	return expression.UnaryExpression.Evaluate()
}