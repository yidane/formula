package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type TernaryExpression struct {
	Left   *opt.LogicalExpression
	Middle *opt.LogicalExpression
	Right  *opt.LogicalExpression
}

func NewTernaryExpression(left, mid, right *opt.LogicalExpression) *opt.LogicalExpression {
	var result opt.LogicalExpression = &TernaryExpression{
		Left:   left,
		Middle: mid,
		Right:  right,
	}

	return &result
}

func (expression *TernaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	left, err := (*expression.Left).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if left.Type != reflect.Bool {
		return nil, fmt.Errorf("ternary need bool first")
	}

	if left.Value.(bool) {
		return (*expression.Middle).Evaluate(context)
	}

	return (*expression.Right).Evaluate(context)
}
