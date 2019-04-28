package opt

type LogicalExpression interface {
	Evaluate(context *FormulaContext) (*Argument, error)
}

type BaseExpression struct {
	Context *FormulaContext
}
