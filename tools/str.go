package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

//零GC 类型转换
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToString(i interface{}) string {
	return fmt.Sprintf("%+v", i)
}

//TO json return string
func TOJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

//TO json return byte
func TOJSONByte(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

//首字幕大写
func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	vv := []byte(str)
	if vv[0] >= 97 && vv[0] <= 122 { //ascii
		vv[0] -= 32
	}
	return string(vv)
}

func Interface2Int64(inter interface{}) (int64, error) {
	switch inter.(type) {
	case string:
		s := inter.(string)
		return strconv.ParseInt(s, 10, 64)
	case int64:
		return inter.(int64), nil
	case int:
		return int64(inter.(int)), nil
	case int32:
		return int64(inter.(int32)), nil
	default:
		return 0, errors.New("unknow type")
	}
}

func Interface2Int(inter interface{}) (int, error) {
	switch inter.(type) {
	case string:
		s := inter.(string)
		return strconv.Atoi(s)
	case int:
		return inter.(int), nil
	case int32:
		return int(inter.(int32)), nil
	case float64:
		return int(inter.(float64)), nil
	default:
		return 0, errors.New("unknow type")
	}
}

func Interface2String(inter interface{}) (string, error) {
	switch inter.(type) {
	case string:
		return inter.(string), nil
	case int:
		return strconv.Itoa(inter.(int)), nil
	default:
		return "", errors.New("unknow type")
	}
}

func StringToUint(args string) uint {
	if args == "" {
		return 0
	}
	r, err := strconv.Atoi(args)
	if err != nil {
		return 0
	}
	return uint(r)
}

func Uint64ToString(args uint64) string {
	return strconv.FormatUint(args, 10)
}

func Int64ToString(args int64) string {
	return strconv.FormatInt(args, 10)
}

func StringToInt(args string) int {
	if args == "" {
		return 0
	}
	r, err := strconv.Atoi(args)
	if err != nil {
		return 0
	}
	return r
}

func StringToInt32(args string) int32 {
	return int32(StringToInt(args))
}

func StringToInt64(args string) (int64, error) {
	return strconv.ParseInt(args, 10, 64)
}

func Float64ToString(f float64) string {
	return fmt.Sprintf("%v", f)
}

func Float32ToString(f float32) string {
	return fmt.Sprintf("%v", f)
}

//StringToFloat64 string到float64
func StringToFloat64(str string) float64 {
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return f
	} else {
		Log.Error(fmt.Sprintf("StringToFloat64Err:%v,%v", err.Error(), str))
		return 0
	}
}

func Int32ToString(i int32) string {
	return fmt.Sprintf("%v", i)
}

//StringToFloat32 string到float32
//func StringToFloat32(str string) float32 {
//	if f, err := strconv.ParseFloat(str, 32); err == nil {
//		return f
//	} else {
//		Log.Error(fmt.Sprintf("StringToFloat64Err:%v,%v", err.Error(), str))
//		return 0
//	}
//}
