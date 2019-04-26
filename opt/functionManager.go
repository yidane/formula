package opt

import "fmt"

var defaultFunctionManager = functionManager{}

type functionManager struct {
}

func Register(f Function) error {
	var i interface{} = f
	if i == nil {
		return fmt.Errorf("arg can not be nil")
	}

	_, ok := i.(BaseFunction)
	if !ok {
		return fmt.Errorf("function should impletment BaseFunction")
	}

	//cache Function

	return nil
}
