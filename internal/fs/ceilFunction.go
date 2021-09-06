package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type CeilFunction struct {
}

func (*CeilFunction) Name() string {
	return "ceil"
}

func (*CeilFunction) AliasName() []string {
	return nil
}

func (f *CeilFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Ceil(v), reflect.Float64), nil
}

func NewCeilFunction() *CeilFunction {
	return &CeilFunction{}
}
