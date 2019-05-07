package fs

import (
	"github.com/yidane/formula/opt"
	"reflect"
)

type PlusFunction struct {
}

func (*PlusFunction) Name() string {
	return "+"
}

func (f *PlusFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.Type == reflect.String && arg1.Type == reflect.String {
		return opt.NewArgumentWithType(arg0.Value.(string)+arg1.Value.(string), arg0.Type), nil
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}
	v1, err := arg1.Float64()
	if err != nil {
		return nil, err
	}
	return opt.NewArgumentWithType(v0+v1, reflect.Float64), nil
}

func NewPlusFunction() *PlusFunction {
	return &PlusFunction{}
}
