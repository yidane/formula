package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
)

type IdentifierExpression struct {
	opt.BaseExpression
	Name string
}

func (expression *IdentifierExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	expression.Context=context
	return nil
}

func (expression *IdentifierExpression) Evaluate() *opt.Argument {
	if expression.Context.Error!=nil{
		return nil
	}

	if p,ok:= expression.Context.Parameters[expression.Name];ok{
		return opt.NewArgument(p)
	}

	expression.Context.Error=fmt.Errorf("variable %s can not be resolved",expression.Name)
	return nil
}

func (*IdentifierExpression) ToString() string {
	panic("implement me")
}

func NewIdentifierExpression(name string) *opt.LogicalExpression{
	var result opt.LogicalExpression= &IdentifierExpression{
		Name:name,
	}

	return &result
}

func NewIdentifier(name string) *IdentifierExpression{
	return &IdentifierExpression{}
}

func NewVarIdentifierExpression(name string) *opt.LogicalExpression{
	var result opt.LogicalExpression= &IdentifierExpression{
		Name:name[1:len(name)-2],
	}

	return &result
}