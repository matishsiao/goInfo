package goInfo

import (
	"fmt"
	"runtime"
)

type GoInfoObject struct {
	GoOS string
	Kernel string
	Core string
	Platform string
	OS string
	Hostname string
	CPUs int
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("GoOS:",gi.GoOS)
	fmt.Println("Kernel:",gi.Kernel)
	fmt.Println("Core:",gi.Core)
	fmt.Println("Platform:",gi.Platform)
	fmt.Println("OS:",gi.OS)
	fmt.Println("Hostname:",gi.Hostname)
	fmt.Println("CPUs:",gi.CPUs)
}

