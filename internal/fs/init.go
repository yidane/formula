package fs

import "github.com/yidane/formula/opt"

func init() {
	fs := []opt.Function{
		AbsFunction{},
		AcosFunction{},
		AsinFunction{},
		AtanFunction{},
		CeilingFunction{},
		ConcatFunction{},
		CosFunction{},
		DivideFunction{},
		ExpFunction{},
		FloorFunction{},
		GreaterFunction{},
		IEEERemainderFunction{},
		IIFFunction{},
		InFunction{},
		Log2Function{},
		Log10Function{},
		LogFunction{},
		MaxFunction{},
		MinFunction{},
		ModuloFunction{},
		MultiplyFunction{},
		PlusFunction{},
		PowerFunction{},
		RoundFunction{},
		SignFunction{},
		SinFunction{},
		SqrtFunction{},
		SubtractFunction{},
		TanFunction{},
		TruncateFunction{},
	}

	for _, f := range fs {
		err := opt.Register(f)
		if err != nil {
			panic(err)
		}
	}
}
