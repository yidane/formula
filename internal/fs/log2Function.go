package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type Log2Function struct {
}

func (*Log2Function) Name() string {
	return "log2"
}

func (*Log2Function) AliasName() []string {
	return nil
}

func (f *Log2Function) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Log2(v), reflect.Float64), nil
}

func NewLog2Function() *Log2Function {
	return &Log2Function{}
}
