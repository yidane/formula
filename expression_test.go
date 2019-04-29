package formula

import (
	"github.com/yidane/formula/opt"
	"math"
	"testing"
)

func TestNewExpression(t *testing.T) {
	const expString = "1+2"

	exp := NewExpression(expString, opt.IgnoreCase)

	result, err := exp.Evaluate()
	if err != nil {
		t.Fatal(err)
	}

	v, err := result.Int64()
	if err != nil {
		t.Fatal(err)
	}

	if v != 3 {
		t.Fatal()
	}
}

func TestCompute(t *testing.T) {
	exp := NewExpression("(1+2)*3/6+3*2")
	result, err := exp.Evaluate()
	if err != nil {
		t.Fatal(err)
	}

	v, err := result.Float64()
	if err != nil {
		t.Fatal(err)
	}

	if v != 7.5 {
		t.Fatal()
	}
}

func TestCommonExpression(t *testing.T) {
	testCases := []struct {
		exp    string
		result float64
	}{
		{"1+2", 3},
		{"1+2-3", 0},
		{"1+2-3*4", -9},
		{"(1+2)*3/6", 1.5},
		{"(1+2)*3/(2+4)", 1.5},
		{"(1+2)*3/(2+4-3+4-1)", 1.5},
		{"(1+2)*3/(2+4-3+4-1*0-1)", 1.5},
		{"(1+2)*3/(2+4-3+4-1*0-1)*0", 0},
	}

	for _, tt := range testCases {
		expression := NewExpression(tt.exp)
		result, err := expression.Evaluate()

		if err != nil {
			t.Fatal(err)
		}

		v, err := result.Float64()
		if err != nil {
			t.Fatal(err)
		}

		if v != tt.result {
			t.Fatal()
		}
	}
}

func TestFunctionExpression(t *testing.T) {
	testCases := []struct {
		exp    string
		result float64
	}{
		{"sin(π/2)", math.Sin(math.Pi / 2)},
		{"cos(π/2)", math.Cos(math.Pi / 2)},
		{"tan(π/2)", math.Tan(math.Pi / 2)},
		{"asin(sin(π/2))", math.Pi / 2},
	}

	for _, tt := range testCases {
		expression := NewExpression(tt.exp)
		result, err := expression.Evaluate()

		if err != nil {
			t.Fatal(err)
		}

		v, err := result.Float64()
		if err != nil {
			t.Fatal(err)
		}

		if v != tt.result {
			t.Fatal()
		}
	}
}

func BenchmarkOnePlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("1+1")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			b.Fatal(err)
		}
		if v != 2 {
			b.Fatal()
		}
	}
}

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression := NewExpression("1")
		result, err := expression.Evaluate()
		if err != nil {
			b.Fatal(err)
		}

		v, err := result.Int64()
		if err != nil {
			b.Fatal(err)
		}
		if v != 1 {
			b.Fatal()
		}
	}
}

func TestImplements(t *testing.T) {
	//通过断言判断类型是否实现接口或组合了其他结构
	//var i interface{} = opt.Function{}
	//f, ok := i.(opt.Function)
	//
	//fmt.Println(ok)
	//fmt.Println(f)
	//
	//var ii interface{} = fs.AddFunction{}
	//f, ok = ii.(opt.Function)
	//fmt.Println(ok)
	//fmt.Println(f.Evaluate())

	//d := importer.For("source", nil)
	//pkg, err := d.Import("github.com/yidane/formula/internal/fs")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//for _, declName := range pkg.Scope().Names() {
	//	t := pkg.Scope().Lookup(declName).Type()
	//}
}
