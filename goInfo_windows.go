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
		gio := GoInfoObject{Kernel: "windows", Core: "unknown", Platform: platform(), OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
		gio.Hostname, _ = os.Hostname()
		return gio, fmt.Errorf("getInfo: %s", err)
	}
	ver := "unknown"
	outStr := out.String()
	if strings.Contains(outStr, "[") {
		parts := strings.Split(strings.ReplaceAll(outStr, "\r\n\t", ""), "[")
		if len(parts) >= 2 {
			x := parts[1]
			x = strings.Split(x, "]")[0]
			parts = strings.Split(x, " ")
			ver = parts[len(parts)-1]
		}
	}
	gio := GoInfoObject{Kernel: "windows", Core: ver, Platform: platform(), OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()
	return gio, nil
}

func platform() string {
	p := runtime.GOARCH
	switch p {
	case "386":
		return "i386"
	case "amd64":
		return "x86_64"
	default:
		return p

	}
}
