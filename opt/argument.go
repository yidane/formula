package opt

import (
	"fmt"
	"reflect"
)

type Argument struct {
	Value interface{}
	Type  reflect.Kind
}

func NewArgument(v interface{}) *Argument {
	if v == nil {
		return &Argument{
			Value: "",
			Type:  reflect.String,
		}
	}

	return &Argument{
		Value: v,
		Type:  reflect.TypeOf(v).Kind(),
	}
}

func NewArgumentWithType(v interface{}, t reflect.Kind) *Argument {
	return &Argument{
		Value: v,
		Type:  t,
	}
}

func (arg *Argument) Int64() (int64, error) {
	switch arg.Type {
	case reflect.Int:
		return int64(arg.Value.(int)), nil
	case reflect.Int8:
		return int64(arg.Value.(int8)), nil
	case reflect.Int16:
		return int64(arg.Value.(int16)), nil
	case reflect.Int32:
		return int64(arg.Value.(int32)), nil
	case reflect.Int64:
		return arg.Value.(int64), nil
	case reflect.Uint:
		return int64(arg.Value.(uint)), nil
	case reflect.Uint8:
		return int64(arg.Value.(uint8)), nil
	case reflect.Uint16:
		return int64(arg.Value.(uint16)), nil
	case reflect.Uint32:
		return int64(arg.Value.(uint32)), nil
	case reflect.Uint64:
		return int64(arg.Value.(uint64)), nil
	case reflect.Float32:
		return int64(arg.Value.(float32)), nil
	case reflect.Float64:
		return int64(arg.Value.(float64)), nil
	default:
		return 0, fmt.Errorf("type %s can not convert to int64", arg.Type)
	}
}

func (arg *Argument) Float64() (float64, error) {
	switch arg.Type {
	case reflect.Int:
		return float64(arg.Value.(int)), nil
	case reflect.Int8:
		return float64(arg.Value.(int)), nil
	case reflect.Int16:
		return float64(arg.Value.(int16)), nil
	case reflect.Int32:
		return float64(arg.Value.(int32)), nil
	case reflect.Int64:
		return float64(arg.Value.(int64)), nil
	case reflect.Uint:
		return float64(arg.Value.(uint)), nil
	case reflect.Uint8:
		return float64(arg.Value.(uint8)), nil
	case reflect.Uint16:
		return float64(arg.Value.(uint16)), nil
	case reflect.Uint32:
		return float64(arg.Value.(uint32)), nil
	case reflect.Uint64:
		return float64(arg.Value.(uint64)), nil
	case reflect.Float32:
		return float64(arg.Value.(float32)), nil
	case reflect.Float64:
		return arg.Value.(float64), nil
	default:
		return 0, fmt.Errorf("type %s can not convert to float64", arg.Type)
	}
}

func (arg *Argument) IsNumber() bool {
	return arg.Type == reflect.Int ||
		arg.Type == reflect.Int8 ||
		arg.Type == reflect.Int16 ||
		arg.Type == reflect.Int32 ||
		arg.Type == reflect.Int64 ||
		arg.Type == reflect.Float32 ||
		arg.Type == reflect.Float64 ||
		arg.Type == reflect.Uint ||
		arg.Type == reflect.Uint8 ||
		arg.Type == reflect.Uint16 ||
		arg.Type == reflect.Uint32 ||
		arg.Type == reflect.Uint64
}
