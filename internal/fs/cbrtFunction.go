package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type CbrtFunction struct {
}

func (*CbrtFunction) Name() string {
	return "cbrt"
}

func (*CbrtFunction) AliasName() []string {
	return nil
}

func (f *CbrtFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Cbrt(v), reflect.Float64), nil
}

func NewCbrtFunction() *CbrtFunction {
	return &CbrtFunction{}
}
