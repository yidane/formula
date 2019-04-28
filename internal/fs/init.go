package fs

import "github.com/yidane/formula/opt"

func init() {
	fs := []opt.Function{
		NewAbsFunction(),
		NewAcosFunction(),
		NewAsinFunction(),
		NewAtanFunction(),
		NewCeilingFunction(),
		NewConcatFunction(),
		NewCosFunction(),
		NewDivideFunction(),
		NewExpFunction(),
		NewFloorFunction(),
		NewGreaterFunction(),
		NewIEEERemainderFunction(),
		NewIIFFunction(),
		NewInFunction(),
		NewLog2Function(),
		NewLog10Function(),
		NewLogFunction(),
		NewMaxFunction(),
		NewMinFunction(),
		NewModuloFunction(),
		NewMultiplyFunction(),
		NewPlusFunction(),
		NewPowerFunction(),
		NewRoundFunction(),
		NewSignFunction(),
		NewSinFunction(),
		NewSqrtFunction(),
		NewSubtractFunction(),
		NewTanFunction(),
		NewTruncateFunction(),
	}

	for _, f := range fs {
		err := opt.Register(&f)
		if err != nil {
			panic(err)
		}
	}
}
