package exp

import (
	"github.com/yidane/formula/opt"
	"reflect"
	"time"
)

type StringValueExpression struct {
	opt.BaseExpression
	Value string
}

func NewStringValueExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &StringValueExpression{
		Value:value,
	}

	return &result
}

func (expression *StringValueExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	expression.Context=context
	return nil
}

func (expression *StringValueExpression) Evaluate() *opt.Argument {
	return opt.NewArgumentWithType(expression.Value,reflect.String)
}

func (*StringValueExpression) ToString() string {
	return ""
}

type IntegerValueExpression struct {
	Value int64
}

func NewIntegerValueExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &IntegerValueExpression{

	}

	return &result
}

func (*IntegerValueExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*IntegerValueExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*IntegerValueExpression) ToString() string {
	panic("implement me")
}

type FloatExpression struct {
	Value float64
}

func NewFloatExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FloatExpression{
	}

	return &result
}

func (*FloatExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*FloatExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*FloatExpression) ToString() string {
	panic("implement me")
}

type DateTimeExpression struct {
	Value time.Time
}

func NewDateTimeExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &DateTimeExpression{}
	return &result
}

func (*DateTimeExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*DateTimeExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*DateTimeExpression) ToString() string {
	panic("implement me")
}

type BooleanValueExpression struct {
	Value bool
}

func NewBooleanValueExpression(value bool) *opt.LogicalExpression {
	 var result opt.LogicalExpression = &BooleanValueExpression{
		Value:value,
	}

	 return &result
}

func (*BooleanValueExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*BooleanValueExpression) Evaluate() *opt.Argument {
	panic("implement me")
}

func (*BooleanValueExpression) ToString() string {
	panic("implement me")
}
