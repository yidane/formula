package fs

import (
	"math"
	"reflect"

	"github.com/yidane/formula/opt"
)

type AcosFunction struct {
}

func (*AcosFunction) Name() string {
	return "acos"
}

func (*AcosFunction) AliasName() []string {
	return nil
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
