package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type CosFunction struct {
}

func (*CosFunction) Name() string {
	return "cos"
}

func (f *CosFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchOneArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Cos(v0), reflect.Float64), nil
}

func NewCosFunction() *CosFunction {
	return &CosFunction{}
}
