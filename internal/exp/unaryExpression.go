package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type UnaryExpression struct {
	Name       string
	Expression *opt.LogicalExpression
}

type NotUnaryExpression struct {
	UnaryExpression
}

func NewNotUnaryExpression(name string, expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &NotUnaryExpression{
		UnaryExpression{
			Name:       name,
			Expression: expression,
		},
	}

	return &result
}

func (expression *NotUnaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	arg, err := (*expression.Expression).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg.Type != reflect.Bool {
		return nil, fmt.Errorf("unary expression ! required bool argument")
	}

	if arg.Value.(bool) {
		return opt.NewArgumentWithType(true, reflect.Bool), nil
	}

	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

type BitwiseNotUnaryExpression struct {
	UnaryExpression
}

func NewBitwiseNotUnaryExpression(name string, expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &BitwiseNotUnaryExpression{
		UnaryExpression{
			Name:       name,
			Expression: expression,
		},
	}

	return &result
}

func (expression *BitwiseNotUnaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	arg, err := (*expression.Expression).Evaluate(context)
	if err != nil {
		return nil, err
	}

	switch arg.Type {
	case reflect.Int:
		return opt.NewArgumentWithType(^arg.Value.(int), reflect.Int), nil
	case reflect.Int8:
		return opt.NewArgumentWithType(^arg.Value.(int8), reflect.Int8), nil
	case reflect.Int16:
		return opt.NewArgumentWithType(^arg.Value.(int16), reflect.Int16), nil
	case reflect.Int32:
		return opt.NewArgumentWithType(^arg.Value.(int32), reflect.Int32), nil
	case reflect.Int64:
		return opt.NewArgumentWithType(^arg.Value.(int64), reflect.Int64), nil
	case reflect.Uint:
		return opt.NewArgumentWithType(^arg.Value.(uint), reflect.Uint), nil
	case reflect.Uint8:
		return opt.NewArgumentWithType(^arg.Value.(uint8), reflect.Uint8), nil
	case reflect.Uint16:
		return opt.NewArgumentWithType(^arg.Value.(uint16), reflect.Uint16), nil
	case reflect.Uint32:
		return opt.NewArgumentWithType(^arg.Value.(uint32), reflect.Uint32), nil
	case reflect.Uint64:
		return opt.NewArgumentWithType(^arg.Value.(uint64), reflect.Uint64), nil
	default:
		return nil, fmt.Errorf("unary expression ! required integer argument")
	}
}

type NegateUnaryExpression struct {
	UnaryExpression
}

func NewNegateUnaryExpression(name string, expression *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &NegateUnaryExpression{
		UnaryExpression{
			Name:       name,
			Expression: expression,
		},
	}

	return &result
}

func (expression *NegateUnaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	arg, err := (*expression.Expression).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if !arg.IsNumber() {
		return nil, fmt.Errorf("unary expression ! required number argument")
	}

	v, _ := arg.Float64()

	return opt.NewArgumentWithType(-v, reflect.Float64), nil
}
