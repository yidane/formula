package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type AbsFunction struct {
}

func (*AbsFunction) Name() string {
	return "Abs"
}

func (f *AbsFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Abs(v), reflect.Float64), nil
}

func NewAbsFunction() *AbsFunction {
	return &AbsFunction{}
}
