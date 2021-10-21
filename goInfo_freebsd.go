package goInfo

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GetInfo() (GoInfoObject, error) {
	out, err := _getInfo()
	for strings.Index(out, "broken pipe") != -1 {
		out, err = _getInfo()
		time.Sleep(500 * time.Millisecond)
	}
	osStr := strings.Replace(out, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	gio := GoInfoObject{Kernel: osInfo[0], Core: osInfo[1], Platform: runtime.GOARCH, OS: osInfo[2], GoARCH: runtime.GOARCH, GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()
	return gio, err
}

func _getInfo() (string, error) {
	cmd := exec.Command("uname", "-sri")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return out.String(), err
}
