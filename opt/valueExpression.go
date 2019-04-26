package opt

import (
	"reflect"
	"time"
)

type StringValueExpression struct {
	BaseExpression
	Value string
}

func NewStringValueExpression(value string) *LogicalExpression{
	var result LogicalExpression= &StringValueExpression{
		Value:value,
	}

	return &result
}

func (expression *StringValueExpression) Accept(context *FormulaContext) *LogicalExpression {
	expression.Context=context
	return nil
}

func (expression *StringValueExpression) Evaluate() *Argument {
	return NewArgumentWithType(expression.Value,reflect.String)
}

func (*StringValueExpression) ToString() string {
	return ""
}

type IntegerValueExpression struct {
	Value int64
}

func NewIntegerValueExpression(value string) *LogicalExpression{
	var result LogicalExpression= &IntegerValueExpression{

	}

	return &result
}

func (*IntegerValueExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*IntegerValueExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*IntegerValueExpression) ToString() string {
	panic("implement me")
}

type FloatExpression struct {
	Value float64
}

func NewFloatExpression(value string) *LogicalExpression{
	var result LogicalExpression= &FloatExpression{
	}

	return &result
}

func (*FloatExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*FloatExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*FloatExpression) ToString() string {
	panic("implement me")
}

type DateTimeExpression struct {
	Value time.Time
}

func NewDateTimeExpression(value string) *LogicalExpression{
	var result LogicalExpression= &DateTimeExpression{}
	return &result
}

func (*DateTimeExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*DateTimeExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*DateTimeExpression) ToString() string {
	panic("implement me")
}

type BooleanValueExpression struct {
	Value bool
}

func NewBooleanValueExpression(value bool) *LogicalExpression{
	 var result LogicalExpression= &BooleanValueExpression{
		Value:value,
	}

	 return &result
}

func (*BooleanValueExpression) Accept(context *FormulaContext) *LogicalExpression {
	panic("implement me")
}

func (*BooleanValueExpression) Evaluate() *Argument {
	panic("implement me")
}

func (*BooleanValueExpression) ToString() string {
	panic("implement me")
}
