package opt

type Function interface {
	Name() string
	LowerName() string
	Evaluate() Argument
}

func MatchArgument(function Function, args ...LogicalExpression) error{
	if args==nil||len(args)==0{
		return nil
	}

	return nil
}

type FunctionManager struct {

}

