//go:build darwin
// +build darwin

package tools

import (
	"github.com/go-cmd/cmd"
)

func GetTickCount() int {
	c := cmd.NewCmd("bash", "-c", "sysctl -a |grep kern.boottime|awk '{print $8}'")
	<-c.Start()
	return StringToInt(c.Status().Stdout[0])
}
