package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type CosFunction struct {
}

func (*CosFunction) Name() string {
	return "cos"
}

func (f *CosFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Cos(v), reflect.Float64), nil
}

func NewCosFunction() *CosFunction {
	return &CosFunction{}
}
