package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type ExpFunction struct {
}

func (*ExpFunction) Name() string {
	return "exp"
}

func (f *ExpFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Exp(v), reflect.Float64), nil
}

func NewExpFunction() *ExpFunction {
	return &ExpFunction{}
}
