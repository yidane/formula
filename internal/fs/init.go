package fs

import (
	"github.com/yidane/formula/internal/cache"
	"github.com/yidane/formula/opt"
)

func init() {
	fs := []opt.Function{
		NewAbsFunction(),
		NewAcosFunction(),
		NewAsinFunction(),
		NewAtanFunction(),
		NewCbrtFunction(),
		NewCeilFunction(),
		NewConcatFunction(),
		NewCosFunction(),
		NewDivideFunction(),
		NewExpFunction(),
		NewFloorFunction(),
		NewGreaterFunction(),
		NewIIFFunction(),
		NewInFunction(),
		NewLessFunction(),
		NewLogFunction(),
		NewLog2Function(),
		NewLog10Function(),
		NewLnFunction(),
		NewMaxFunction(),
		NewMinFunction(),
		NewModFunction(),
		NewMultiplyFunction(),
		NewPlusFunction(),
		NewPowerFunction(),
		NewRoundFunction(),
		NewShiftLeftFunction(),
		NewShiftRightFunction(),
		NewSignFunction(),
		NewSinFunction(),
		NewSqrtFunction(),
		NewSubtractFunction(),
		NewTanFunction(),
		NewTruncateFunction(),
	}

	for i := 0; i < len(fs); i++ {
		err := cache.Register(&fs[i])
		if err != nil {
			panic(err)
		}
	}
}

func ParseFloat(name string, context *opt.FormulaContext, args ...*opt.LogicalExpression) (float64, error) {
	err := opt.MatchOneArgument(name, args...)
	if err != nil {
		return 0, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return 0, err
	}

	v, err := arg0.Float64()
	if err != nil {
		return 0, err
	}

	return v, nil
}
