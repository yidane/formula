package opt

import (
	"fmt"
)

type Function interface {
	Evaluate(args ...LogicalExpression) *Argument
}

type BaseFunction struct {
	context   FormulaContext
	Name      func() string
	LowerName func() string
}

func (f BaseFunction) SetContext(ctx FormulaContext) {
	f.context = ctx
}

func (f BaseFunction) GetContext() FormulaContext {
	return f.context
}

func (f BaseFunction) ReportError(err error) {
	f.context.Error = err
}

func (f BaseFunction) HasError() bool{
	return f.context.Error!=nil
}

func (f BaseFunction) MatchArgument(args ...LogicalExpression) {
	if f.context.Error != nil {
		return
	}

	if args == nil || len(args) == 0 {
		f.context.Error = fmt.Errorf(f.Name())
	}
}
