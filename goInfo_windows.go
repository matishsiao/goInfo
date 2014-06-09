package goInfo

import (		
	"strings"	
	"os/exec"
	"os"
	"bytes"
)

func GetInfo() *GoInfoObject {			
	cmd := exec.Command("cmd","SystemInfo")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {							
		panic(err)	
	}
	fmt.Println("Win:",out.String())	
	osStr := strings.Replace(out.String(),"\n","",-1)
	osStr = strings.Replace(osStr,"\r\n","",-1)
	osInfo := strings.Split(osStr," ")
	gio := new(GoInfoObject)
	gio.Kernel = osInfo[0]
	gio.Core = osInfo[1]
	gio.Platform = osInfo[2]
	gio.OS = osInfo[3]
	gio.Hostname,_ = os.Hostname()	
	return gio
}