package goInfo

import (		
	"strings"	
	"os/exec"
	"bytes"
)

func GetInfo() *GoInfoObject {			
	cmd := exec.Command("uname","-srio")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {							
		panic(err)	
	}
	//Kernel hostname version machine processor hardware os
	//Linux ubuntu-Matis-VM 3.13.0-24-generic #47-Ubuntu SMP Fri May 2 23:30:00 UTC 2014 x86_64 x86_64 x86_64 GNU/Linux	
	osInfo := strings.Split(out.String()," ")
	gio := new(GoInfoObject)
	gio.Kernel = osInfo[0]
	gio.Core = osInfo[1]
	gio.Platform = osInfo[2]
	gio.OS = osInfo[3]
	return gio
}