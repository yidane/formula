package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type Log10Function struct {
}

func (*Log10Function) Name() string {
	return "log10"
}

func (f *Log10Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
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

	return opt.NewArgumentWithType(math.Log10(v0), reflect.Float64), nil
}

func NewLog10Function() *Log10Function {
	return &Log10Function{}
}
