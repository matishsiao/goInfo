package goInfo

import (		
	"strings"	
	"os/exec"
	"os"
	"bytes"
	"runtime"
)

func GetInfo() *GoInfoObject {			
	cmd := exec.Command("cmd","ver")
	cmd.Stdin = strings.NewReader("some input")	
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {							
		panic(err)	
	}	
	osStr := strings.Replace(out.String(),"\n","",-1)
	osStr = strings.Replace(osStr,"\r\n","",-1)
	tmp1 := strings.Index(osStr,"[Version")
	tmp2 := strings.Index(osStr,"]")
	var ver string
	if tmp1 == -1 || tmp2 == -1 {
		ver = "unknown"
	} else {
		ver = osStr[tmp1+9:tmp2]
	}
	
	gio := new(GoInfoObject)
	gio.GoOS = runtime.GOOS
	gio.Kernel = "windows"
	gio.Core = ver
	gio.Platform = "unknown"
	gio.OS = "windows"
	gio.CPUs = runtime.NumCPU()
	gio.Hostname,_ = os.Hostname()	
	return gio
}