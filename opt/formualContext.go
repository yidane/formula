package opt

import "bytes"

type Options int

const (
	None=Options(1)
	IgnoreCase=Options(2)
	NoCache=Options(4)
	IterateParameters=Options(8)
	RoundAwayFromZero=Options(16)
)

type FormulaContext struct {
	Options Options
	Parameters map[string]interface{}
	Result bytes.Buffer
	Error error
}

func NewFormulaContext(options Options) FormulaContext{
	return FormulaContext{
		Options:options,
		Parameters:make(map[string]interface{}),
	}
}