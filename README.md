# formula

支持对数学运算的编译执行，主要支持以下数学公式：

* 支持四则运算
* 支持三角函数
* 支持位运算
* 支持三元操作
* 支持自定义参数
* 支持自定义函数
* 支持公式缓存

内置函数如下表

| 函数名称 | 参数个数 | 调用方式            | 结果                     |
| -------- | -------- | ------------------- | ------------------------ |
| abs      | 1        | abs(-1)             | 1                        |
| acos     | 1        | acos(sqrt(3)/2)     | 0.5235987755982991 (π/6) |
| asin     | 1        | asin(1/2)           | 0.5235987755982991 (π/6) |
| asin     | 1        | asin(1)             | π/2                      |
| atan     | 1        | atan(1)             | 0.7853981633974483 (π/4) |
| ceil     | 1        | ceil(3.4)           | 4                        |
| concat   | n        | concat(1,23,hello)  | 123hello                 |
| cos      | 1        | cos(π/3)            | 0.5000000000000001       |
| /        | 2        | 3/4                 | 0.75                     |
| exp      | 1        | exp(3.3)            | 27.112638920657883       |
| floor    | 1        | floor(2.2)          | 2                        |
| \>       | 2        | 3 > 2               | true                     |
| iif      | 3        | iif(3 > 2,π,10)     | 3.141592653589793        |
| \<       | 2        | 3 < 2               | false                    |
| in       | n        | in(3,3,4,5)         | true                     |
| ln       | 1        | ln(2.718281828)     | 0.9999999998311266       |
| log2     | 1        | log2(16)            | 4                        |
| log10    | 1        | log10(100000)       | 5                        |
| log      | 2        | log(100,10)         | 2                        |
| max      | n        | max(-1,2,3.1) = 3.1 | 3.1                      |
| min      | n        | min(-1,2,3.1) = -1  | -1                       |
| mod      | 2        | mod(5,2)            | 1                        |
| *        | 2        | 3*3.4               | 10.2                     |
| +        | 2        | 5+10                | 15                       |
| pow      | 2        | pow(10,2)           | 100                      |
| round    | 2        | round(100.11)       | 100                      |
| sign     | 1        | sign(100)           | false                    |
| sin      | 1        | sin(π/6)            | 0.49999999999999994      |
| -        | 2        | 3-6                 | -3                       |
| tan      | 1        | tan(π/4)            | 1                        |
| truncate | 1        | truncate(12.3)      | 12                       |
| \>>       | 2        | 2>>1                | 1                        |
| \<<       | 2        | 1<<1                | 2                        |

## 调用方法

* 普通数学公式调用

``` go

expression:=formula.NewExpression("1+2")
result,err:=expression.Evaluate()
if err!=nil{
    //handle err
}

v,err:= result.Int64()
if err!=nil{
    //handle err
}

//v should equal 3

```

* 自定义参数调用
```go
expression := NewExpression("[i]+[j]")
err := expression.AddParameter("i", 1)
if err != nil {
    t.Fatal(err)
}
err = expression.AddParameter("j", 2)
if err != nil {
    t.Fatal(err)
}
result, err := expression.Evaluate()
//handle result
```

自定义公式开发

* 实现接口 opt.Function
```go
type CustomFunction struct {
}

func (*CustomFunction) Name() string {
	return "CustomFunction"
}

func (f *CustomFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	left, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	leftValue, err := left.Int64()
	if err != nil {
		return nil, err
	}

	right, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	rightValue, err := right.Int64()
	if err != nil {
		return nil, err
	}

	return opt.NewArgumentWithType(leftValue+rightValue+1, reflect.Int64), nil
}
```

* 注册自定义函数
```go
func init() {
	var f opt.Function = new(CustomFunction)
	err := formula.Register(&f)
	if err != nil {
		log.Fatal(err)
	}
}
```

* 调用自定义函数

```go
expression := formula.NewExpression("CustomFunction(1,2)")
result, err := expression.Evaluate()
if err != nil {
    log.Fatal(err)
}

v, err := result.Int64()
if err != nil {
    log.Fatal(err)
}

if v != 4 { //CustomFunction: i+j+1
    log.Fatal("error")
}

log.Println("custom function succeed")
```

## 性能测试

```bash
BenchmarkOnePlusOne-8   	   50000	     26676 ns/op
BenchmarkOne-8   	           100000	     19401 ns/op
BenchmarkComplexOne-8   	   10000	    180650 ns/op
BenchmarkSin-8   	           20000	     78591 ns/op
```