package opt

type LogicalExpression interface {
	Accept(context *FormulaContext) *LogicalExpression
	Evaluate() *Argument
	ToString() string
}

type BaseExpression struct {
	Context *FormulaContext
}

func (expression BaseExpression) ReportErr(err error) {
	expression.Context.Error = err
}

func (expression BaseExpression) LastError() error {
	return expression.Context.Error
}
