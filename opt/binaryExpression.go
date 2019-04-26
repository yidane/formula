package opt

type BinaryExpression struct {
	BaseExpression
	Name            string
	LeftExpression  *LogicalExpression
	RightExpression *LogicalExpression
}

func NewBinaryExpression(name string,leftExpression ,rightExpression *LogicalExpression) *LogicalExpression{
	var result LogicalExpression=  &BinaryExpression{
		Name:name,
		LeftExpression:leftExpression,
		RightExpression:rightExpression,
	}

	return &result
}

func (expression *BinaryExpression) Accept(context *FormulaContext) *LogicalExpression {
	expression.Context=context

	//p:= unsafe.Pointer(expression)

	return nil
}

func (expression *BinaryExpression) Evaluate() *Argument {
	return NewArgument(nil)
}

func (expression *BinaryExpression) ToString() string {
	return ""
}