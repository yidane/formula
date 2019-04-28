package fs

import "github.com/yidane/formula/opt"

type FloorFunction struct {
}

func (*FloorFunction) Name() string {
	panic("implement me")
}

func (*FloorFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewFloorFunction() *FloorFunction {
	return &FloorFunction{}
}
