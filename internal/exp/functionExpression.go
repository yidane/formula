package exp

import (
	"github.com/yidane/formula/internal/cache"
	"github.com/yidane/formula/opt"
)

type FunctionExpression struct {
	BaseExpression
	Identifier *IdentifierExpression
	Arguments  []*opt.LogicalExpression
}

func NewFunctionExpression(id *IdentifierExpression, args []*opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FunctionExpression{
		Identifier: id,
		Arguments:  args,
	}

	return &result
}

func (expression *FunctionExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	f, err := cache.FindFunction(expression.Identifier.Name)
	if err != nil {
		return nil, err
	}

	return (*f).Evaluate(context, expression.Arguments...)
}
