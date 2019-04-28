package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type AsinFunction struct {
}

func (*AsinFunction) Name() string {
	return "asin"
}

func (f *AsinFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchOneArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Asin(v), reflect.Float64), nil
}

func NewAsinFunction() *AsinFunction {
	return &AsinFunction{}
}
