package exp

import "github.com/yidane/formula/opt"

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

func (*TernaryExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	panic("implement me")
}
