package opt

import "bytes"

type Option int

const (
	None              = Option(1)
	IgnoreCase        = Option(2)
	NoCache           = Option(4)
	IterateParameters = Option(8)
	RoundAwayFromZero = Option(16)
)

type FormulaContext struct {
	Option     Option
	Parameters map[string]interface{}
	Result     bytes.Buffer
}

func mergeOption(options ...Option) Option {
	if len(options) == 0 {
		return None
	}

	option := int(options[0])

	for i := 0; i < len(options)-1; i++ {
		option = option | int(options[i])
	}

	return Option(option)
}

func NewFormulaContext(options ...Option) *FormulaContext {
	return &FormulaContext{
		Option:     mergeOption(options...),
		Parameters: make(map[string]interface{}),
	}
}
