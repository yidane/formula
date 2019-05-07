package fs

import (
	"github.com/yidane/formula/opt"
)

type PowerFunction struct {
}

func (*PowerFunction) Name() string {
	return "pow"
}

func (*PowerFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	return nil, nil
}

func NewPowerFunction() *PowerFunction {
	return &PowerFunction{}
}
