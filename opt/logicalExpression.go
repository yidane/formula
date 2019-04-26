package opt

import "fmt"

type LogicalExpression interface {
	Accept(context *FormulaContext) *LogicalExpression
	Evaluate() *Argument
	ToString() string
}

type BaseExpression struct {
	Context *FormulaContext
}


type IdentifierExpression struct {
	BaseExpression
	Name string
}

func (expression *IdentifierExpression) Accept(context *FormulaContext) *LogicalExpression {
	expression.Context=context
	return nil
}

func (expression *IdentifierExpression) Evaluate() *Argument {
	if expression.Context.Error!=nil{
		return nil
	}

	if p,ok:= expression.Context.Parameters[expression.Name];ok{
		return NewArgument(p)
	}

	expression.Context.Error=fmt.Errorf("variable %s can not be resolved",expression.Name)
	return nil
}

func (*IdentifierExpression) ToString() string {
	panic("implement me")
}

func NewIdentifierExpression(name string) *LogicalExpression{
	var result LogicalExpression= &IdentifierExpression{
		Name:name,
	}

	return &result
}

func NewIdentifier(name string) *IdentifierExpression{
	return &IdentifierExpression{}
}

func NewVarIdentifierExpression(name string) *LogicalExpression{
	var result LogicalExpression= &IdentifierExpression{
		Name:name[1:len(name)-2],
	}

	return &result
}