# formula
支持对数学运算的编译执行，主要支持以下数学公式：
* 支持四则运算
* 支持三角函数
* 支持位运算
* 支持三元操作
* 支持自定义参数
* 支持自定义函数
* 支持公式缓存

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
