package opt

type LogicalExpression interface {
	Evaluate(context *FormulaContext) (*Argument, error)
}
