package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type AsinFunction struct {
}

func (*AsinFunction) Name() string {
	return "asin"
}

func (*AsinFunction) AliasName() []string {
	return nil
}

func (f *AsinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Asin(v), reflect.Float64), nil
}

func NewAsinFunction() *AsinFunction {
	return &AsinFunction{}
}
