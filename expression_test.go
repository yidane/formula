package formula

import (
	"github.com/yidane/formula/opt"
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
