package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type FloorFunction struct {
}

func (*FloorFunction) Name() string {
	return "floor"
}

func (f *FloorFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Floor(v), reflect.Float64), nil
}

func NewFloorFunction() *FloorFunction {
	return &FloorFunction{}
}
