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

	return opt.NewArgumentWithType(math.Log(v0)/math.Log(v1), reflect.Float64), nil
}

func NewLogFunction() *LogFunction {
	return &LogFunction{}
}
