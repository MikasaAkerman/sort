package sort

import (
	"reflect"
	"sort"
)

type OrderType int

const (
	Asc OrderType = iota
	Desc
)

type Option struct {
	Fields    string // 支持 float64/float32/string/int64/int/int32/int16/int8 类型字段排序
	OrderType OrderType
}

func Sort(list interface{}, ops ...Option) {
	if list == nil {
		return
	}
	if len(ops) == 0 {
		return
	}
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice {
		return
	}
	size := v.Len()
	switch size {
	case 0, 1:
		return
	}

	sort.Slice(list, func(i, j int) bool {
		for _, op := range ops {
			if len(op.Fields) == 0 {
				continue
			}
			var f1 reflect.Value
			var f2 reflect.Value
			switch v.Index(i).Kind() {
			case reflect.Interface, reflect.Ptr:
				f1 = v.Index(i).Elem().FieldByName(op.Fields)
				f2 = v.Index(j).Elem().FieldByName(op.Fields)
			default:
				f1 = v.Index(i).FieldByName(op.Fields)
				f2 = v.Index(j).FieldByName(op.Fields)
			}

			switch f1.Kind() {
			case reflect.Float64, reflect.Float32:
				if f1.Float() > f2.Float() {
					//if op.OrderType == Desc {
					//	return true
					//} else {
					//	return false
					//}
					return op.OrderType == Desc
				}
				if f1.Float() < f2.Float() {
					return op.OrderType == Asc
				}
			case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8:
				if f1.Int() > f2.Int() {
					return op.OrderType == Desc
				}
				if f1.Int() < f2.Int() {
					return op.OrderType == Asc
				}

			case reflect.String:
				if f1.String() > f2.String() {
					return op.OrderType == Desc
				}
				if f1.String() < f2.String() {
					return op.OrderType == Asc
				}
			}
		}
		return false
	})
}

type Func func(i, j int) int // -1 小于，0 等于, 1 大于

func ByFunc(list interface{}, fns ...Func) {
	sort.Slice(list, func(i, j int) bool {
		for _, fn := range fns {
			r := fn(i, j)
			switch {
			case r < 0:
				return true
			case r > 0:
				return false
			case r == 0: // 值相等，根据下一个排序方法排序
			}
		}
		return false
	})
}
