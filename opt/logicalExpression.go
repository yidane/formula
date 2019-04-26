package opt

type LogicalExpression interface {
	Accept(context *FormulaContext) *LogicalExpression
	Evaluate() *Argument
	ToString() string
}

type BaseExpression struct {
	Context *FormulaContext
}