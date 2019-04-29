package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type LogFunction struct {
}

func (*LogFunction) Name() string {
	return "log"
}

func (f *LogFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Log(v), reflect.Float64), nil
}

func NewLogFunction() *LogFunction {
	return &LogFunction{}
}
