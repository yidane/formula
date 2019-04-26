package opt

import "reflect"

type Argument struct {
	Value interface{}
	Type reflect.Kind
}

func NewArgument(v interface{}) *Argument{
	if v==nil{
		return &Argument{
			Value:"",
			Type:reflect.String,
		}
	}

	return &Argument{
		Value:v,
		Type:reflect.TypeOf(v).Kind(),
	}
}

func NewArgumentWithType(v interface{},t reflect.Kind) *Argument{
	return &Argument{
		Value:v,
		Type:t,
	}
}

func (arg *Argument) IsNumber() bool{
	return arg.Type==reflect.Int||
		arg.Type==reflect.Int8||
		arg.Type==reflect.Int16||
		arg.Type==reflect.Int32||
		arg.Type==reflect.Int64||
		arg.Type==reflect.Float32||
		arg.Type==reflect.Float64||
		arg.Type==reflect.Uint||
		arg.Type==reflect.Uint8||
		arg.Type==reflect.Uint16||
		arg.Type==reflect.Uint32||
		arg.Type==reflect.Uint64
}
