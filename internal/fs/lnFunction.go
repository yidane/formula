package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type LnFunction struct {
}

func (*LnFunction) Name() string {
	return "ln"
}

func (*LnFunction) AliasName() []string {
	return nil
}

func (f *LnFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Log(v), reflect.Float64), nil
}

func NewLnFunction() *LnFunction {
	return &LnFunction{}
}
