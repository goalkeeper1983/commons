//go:build windows
// +build windows

package tools

import (
	"fmt"
	"syscall"
	"unsafe"
)

func GetTickCount() int {
	user32 := syscall.MustLoadDLL("kernel32.dll")
	getTickCount := user32.MustFindProc("GetTickCount")
	fmt.Printf("%T\n", getTickCount)
	fmt.Println(getTickCount)
	fmt.Println(getTickCount.Name)
	ticks, _, err := getTickCount.Call()
	fmt.Println(err)
	p := (*int)(unsafe.Pointer(&ticks))
	Log.Info(fmt.Sprintf("%v", *p))
	return *p
}
