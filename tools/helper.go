package tools

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func Catch() {
	if r := recover(); r != nil {
		Log.Error(fmt.Sprintf("%v", r)) //输出panic信息
		Log.Error(string(debug.Stack()[:]))
	}
}

//[a,b]
func GetRandInt(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return a + rand.Intn(b-a+1)
}

var randStringPool = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "N", "M", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "m", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "6", "7", "8", "9"}

func GetRandString() string {
	count := len(randStringPool)
	nickName := ""
	for i := 0; i != 6; i++ {
		nickName = nickName + randStringPool[GetRandInt(0, count-1)]
	}
	return nickName
}
