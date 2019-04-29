package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type TruncateFunction struct {
}

func (*TruncateFunction) Name() string {
	return "truncate"
}

func (f *TruncateFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Trunc(v), reflect.Float64), nil
}

func NewTruncateFunction() *TruncateFunction {
	return &TruncateFunction{}
}
