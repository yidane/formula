package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type SinFunction struct {
}

func (*SinFunction) Name() string {
	return "sin"
}

func (f *SinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Sin(v), reflect.Float64), nil
}

func NewSinFunction() *SinFunction {
	return &SinFunction{}
}
