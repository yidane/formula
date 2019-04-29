package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type RoundFunction struct {
}

func (*RoundFunction) Name() string {
	return "round"
}

func (f *RoundFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Round(v), reflect.Float64), nil
}

func NewRoundFunction() *RoundFunction {
	return &RoundFunction{}
}
