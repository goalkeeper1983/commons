package functionRegister

import (
	"reflect"
	"strings"

	"commons/tools"

	"go.uber.org/zap"
)

var ModelFuncMap = make(map[uint16]reflect.Value)

func GetModelFunc(MsgId uint16) (reflect.Value, bool) {
	if mFunc, ok := ModelFuncMap[MsgId]; ok {
		return mFunc, true
	}
	return reflect.ValueOf(nil), false
}

//每个model的注册方法
func RegisterWorkModelByMsgId(mode interface{}) {
	vf := reflect.ValueOf(mode)
	tf := reflect.TypeOf(mode)
	tools.Log.Info("RegisterWorkModelByMsgId", zap.Any("ValueOf", vf), zap.Any("TypeOf", tf))
	for i := 0; i < tf.NumMethod(); i++ { //取方法
		functionName := tf.Method(i).Name
		countSplit := strings.Split(functionName, "_")
		if len(countSplit) == 2 {
			if cmdId := tools.StringToInt(countSplit[1]); cmdId != 0 {
				tools.Log.Info("RegisterFunction:", zap.Any("MsgId:"+countSplit[1], tf.Method(i)))
				ModelFuncMap[uint16(cmdId)] = vf.MethodByName(functionName)
			}
		}
	}
}

func RegisterFunc(inter interface{}) {
	vf := reflect.ValueOf(inter)
	tf := reflect.TypeOf(inter)
	tools.Log.Info("RegisterFunc", zap.Any("ValueOf", vf), zap.Any("TypeOf", tf))
	for i := 0; i < tf.NumMethod(); i++ { //取方法
		functionName := tf.Method(i).Name
		tools.Log.Info(functionName, zap.Any("Method:", tf.Method(i)), zap.Any("MethodByName", vf.MethodByName(functionName)))
	}
}
