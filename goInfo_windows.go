package goInfo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetInfo() (GoInfoObject, error) {
	cmd := exec.Command("cmd", "ver")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		gio := GoInfoObject{Kernel: "windows", Core: "unknown", Platform: "unknown", OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
		gio.Hostname, _ = os.Hostname()
		return gio, fmt.Errorf("getInfo: %s", err)
	}
	osStr := strings.Replace(out.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	containsBrackVersion := strings.Contains(osStr, "[Version")
	containsClosingBracket := strings.Contains(osStr, "]")
	var ver string
	if !containsBrackVersion || !containsClosingBracket {
		ver = "unknown"
	} else {
		ver = osStr[containsBrackVersion+9 : containsClosingBracket]
	}
	gio := GoInfoObject{Kernel: "windows", Core: ver, Platform: "unknown", OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()
	return gio, nil
}
