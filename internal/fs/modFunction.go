package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type ModFunction struct {
}

func (*ModFunction) Name() string {
	return "mod"
}

func (f *ModFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Mod(v, 1), reflect.Float64), nil
}

func NewModuloFunction() *ModFunction {
	return &ModFunction{}
}
