package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
)

type IdentifierExpression struct {
	opt.BaseExpression
	Name string
}

func (expression *IdentifierExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	if p, ok := expression.Context.Parameters[expression.Name]; ok {
		return opt.NewArgument(p), nil
	}

	return nil, fmt.Errorf("variable %s can not be resolved", expression.Name)
}

func NewIdentifierExpression(name string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &IdentifierExpression{
		Name: name,
	}

	return &result
}

func NewIdentifier(name string) *IdentifierExpression {
	return &IdentifierExpression{}
}

//NewVarIdentifierExpression create new custom parameter which output like '[Parameter]'
func NewVarIdentifierExpression(name string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &IdentifierExpression{
		Name: name[1 : len(name)-2],
	}

	return &result
}
