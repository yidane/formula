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
		NewCeilFunction(),
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

	for i := 0; i < len(fs); i++ {
		err := cache.Register(&fs[i])
		if err != nil {
			panic(err)
		}
	}
}
