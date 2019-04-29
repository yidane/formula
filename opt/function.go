package opt

import (
	"fmt"
)

type Function interface {
	Name() string
	Evaluate(context *FormulaContext, args ...*LogicalExpression) (*Argument, error)
}

func MatchArgument(name string, args ...*LogicalExpression) error {
	if args == nil || len(args) == 0 {
		return fmt.Errorf(name)
	}

	return nil
}

func MatchOneArgument(name string, args ...*LogicalExpression) error {
	if len(args) != 1 {
		return fmt.Errorf("function %s required only one argument", name)
	}

	return nil
}

func MatchTwoArgument(name string, args ...*LogicalExpression) error {
	if len(args) != 2 {
		return fmt.Errorf("function %s required only two arguments", name)
	}

	return nil
}
