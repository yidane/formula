package fs

import (
	"github.com/yidane/formula/opt"
	"math"
	"reflect"
)

type AcosFunction struct {
}

func (*AcosFunction) Name() string {
	return "acos"
}

func (f *AcosFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(math.Acos(v), reflect.Float64), nil
}

func NewAcosFunction() *AcosFunction {
	return &AcosFunction{}
}
