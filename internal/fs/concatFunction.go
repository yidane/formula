package fs

import (
	"reflect"
	"strings"

	"github.com/yidane/formula/opt"
)

type ConcatFunction struct {
}

func (*ConcatFunction) Name() string {
	return "concat"
}

func (*ConcatFunction) AliasName() []string {
	return nil
}

func (*ConcatFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) == 0 {
		return opt.NewArgumentWithType("", reflect.String), nil
	}

	buf := strings.Builder{}
	for i := 0; i < len(args); i++ {
		arg, err := (*args[i]).Evaluate(context)
		if err != nil {
			buf.Reset()
			return nil, err
		}

		buf.WriteString(arg.String())
	}

	return opt.NewArgumentWithType(buf.String(), reflect.String), nil
}

func NewConcatFunction() *ConcatFunction {
	return &ConcatFunction{}
}
