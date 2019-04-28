package fs

import "github.com/yidane/formula/opt"

type AcosFunction struct {
}

func (*AcosFunction) Name() string {
	panic("implement me")
}

func (*AcosFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewAcosFunction() *AcosFunction {
	return &AcosFunction{}
}
