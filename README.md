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
| concat   | 1        | concat(1,23,hello)  | 123hello                 |
| cos      | 1        | cos(π/3)            | 0.5000000000000001       |
| /        | 1        | 3/4                 | 0.75                     |
| exp      | 1        | asin(1)             | π/2                      |
| floor    | 1        | floor(2.2)          | 2                        |
| \>       | 1        | 3 > 2               | true                     |
| iif      | 1        | iif(3 > 2,π,10)     | 3.141592653589793        |
| \<       | 1        | 3 < 2               | false                    |
| in       | 1        | asin(1)             | π/2                      |
| log2     | 1        | log2(16)            | 4                        |
| log10    | 1        | log10(100000)       | 5                        |
| log      | 1        | asin(1)             | π/2                      |
| max      | 1        | max(-1,2,3.1) = 3.1 | 3.1                      |
| min      | 1        | min(-1,2,3.1) = -1  | -1                       |
| mod      | 1        | mod(3)              | 0                        |
| *        | 1        | 3*3.4               | 10.2                     |
| +        | 1        | 5+10                | 15                       |
| pow      | 1        | asin(1)             | π/2                      |
| round    | 1        | round(100.11)       | 100                      |
| sign     | 1        | sign(100)           | false                    |
| sin      | 1        | sin(π/6)            | 0.49999999999999994      |
| -        | 1        | 3-6                 | -3                       |
| tan      | 1        | tan(π/4)            | 1                        |
| truncate | 1        | truncate(12.3)      | 12                       |

## 调用方法

普通数学公式调用

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

自定义公式开发

[](examples/custom-function/myAdd.go)

## 性能测试
