package nError

import (
	"fmt"
)

// 错误常量
const (
	ErrorDbConnectFail = 1
)

//  内部业务错误处理
type NError struct {
	errno  int32
	errmsg string
}

// 返回合成错误信息
func (e *NError) Error() string {
	return string(fmt.Sprintf("errno:[%v], errmsg:[%v]", e.errno, e.errmsg))
}

// 返回错误号码
func (e *NError) ErrorByErrno() string {
	return string(e.errno)
}

// 返回错误信息
func (e *NError) ErrorByErrmsg() string {
	return string(e.errmsg)
}

// 直接报错，异常的形式
// 注意当前函数会中断
func Panic(errno int32, msg ...string) {
	errSt := &NError{errno: errno}
	if len(msg) > 0 {
		errSt.errmsg = msg[0]
	}
	panic(errSt)
}
