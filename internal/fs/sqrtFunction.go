package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type SqrtFunction struct {
}

func (*SqrtFunction) Name() string {
	return "sqrt"
}

func (f *SqrtFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Sqrt(v), reflect.Float64), nil
}

func NewSqrtFunction() *SqrtFunction {
	return &SqrtFunction{}
}
