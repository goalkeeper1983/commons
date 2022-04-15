//go:build linux
// +build linux

package tools

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetTickCount() int {
	cmd := exec.Command("cat", "/prco/uptime")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	uptimeStr := strings.Split(string(opBytes), ".")
	uptime, err := strconv.Atoi(uptimeStr[0])
	if err != nil {
		log.Fatal(err)
	}
	return uptime
}
