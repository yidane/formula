package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type ExpFunction struct {
}

func (*ExpFunction) Name() string {
	return "exp"
}

func (*ExpFunction) AliasName() []string {
	return nil
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
