# formula

* 支持四则运算
* 支持三角函数
* 支持位运算
* 支持三元操作
* 支持自定义函数

## 调用方法

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