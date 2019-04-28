package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type Log2Function struct {
}

func (*Log2Function) Name() string {
	return "log2"
}

func (f *Log2Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
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

	return opt.NewArgumentWithType(math.Log2(v0), reflect.Float64), nil
}

func NewLog2Function() *Log2Function {
	return &Log2Function{}
}
