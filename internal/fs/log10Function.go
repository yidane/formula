package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type Log10Function struct {
}

func (*Log10Function) Name() string {
	return "log10"
}

func (*Log10Function) AliasName() []string {
	return nil
}

func (f *Log10Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Log10(v), reflect.Float64), nil
}

func NewLog10Function() *Log10Function {
	return &Log10Function{}
}
