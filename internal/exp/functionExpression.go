package exp

import (
	"github.com/yidane/formula/internal/cache"
	"github.com/yidane/formula/opt"
)

type FunctionExpression struct {
	BaseExpression
	Identifier *opt.LogicalExpression
	Arguments  []*opt.LogicalExpression
}

func NewFunctionExpression(id *opt.LogicalExpression, args []*opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FunctionExpression{
		Identifier: id,
		Arguments:  args,
	}

	return &result
}

func (expression *FunctionExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	name, _ := (*expression.Identifier).Evaluate(context)
	f, err := cache.FindFunction(name.String())
	if err != nil {
		return nil, err
	}

	return (*f).Evaluate(context, expression.Arguments...)
}
