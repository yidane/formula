package fs

import "github.com/yidane/formula/opt"

type AbsFunction struct {
}

func (*AbsFunction) Name() string {
	return "Abs"
}

func (f *AbsFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewAbsFunction() *AbsFunction {
	return &AbsFunction{}
}
