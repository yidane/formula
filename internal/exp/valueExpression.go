package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
	"strconv"
	"time"
)

type StringValueExpression struct {
	opt.BaseExpression
	Value string
}

func NewStringValueExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &StringValueExpression{
		Value: value,
	}

	return &result
}

func (expression *StringValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.String), nil
}

func (*StringValueExpression) ToString() string {
	return ""
}

type IntegerValueExpression struct {
	Value int64
}

func NewIntegerValueExpression(value string) *opt.LogicalExpression {
	i, err := strconv.ParseInt(value, 10, 10)
	if err != nil {
		fmt.Println(err)
	}

	var result opt.LogicalExpression = &IntegerValueExpression{
		Value: i,
	}

	return &result
}

func (expression *IntegerValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Value, reflect.Int64), nil
}

type FloatExpression struct {
	Value float64
}

func NewFloatExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &FloatExpression{}

	return &result
}

func (*FloatExpression) Accept(context *opt.FormulaContext) *opt.LogicalExpression {
	panic("implement me")
}

func (*FloatExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	panic("implement me")
}

type DateTimeExpression struct {
	Value time.Time
}

func NewDateTimeExpression(value string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &DateTimeExpression{}
	return &result
}

func (*DateTimeExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	panic("implement me")
}

type BooleanValueExpression struct {
	Value bool
}

func NewBooleanValueExpression(value bool) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BooleanValueExpression{
		Value: value,
	}

	return &result
}

func (*BooleanValueExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	panic("implement me")
}
