package goInfo

import (
	"fmt"
)

type GoInfoObject struct {
	GoARCH   string
	GoOS     string
	Kernel   string
	Core     string
	Platform string
	OS       string
	Hostname string
	CPUs     int
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("GoARCH:", gi.GoARCH)
	fmt.Println("GoOS:", gi.GoOS)
	fmt.Println("Kernel:", gi.Kernel)
	fmt.Println("Core:", gi.Core)
	fmt.Println("Platform:", gi.Platform)
	fmt.Println("OS:", gi.OS)
	fmt.Println("Hostname:", gi.Hostname)
	fmt.Println("CPUs:", gi.CPUs)
}

func (gi *GoInfoObject) String() string {
	return fmt.Sprintf("GoARCH:%s,GoOS:%v,Kernel:%v,Core:%v,Platform:%v,OS:%v,Hostname:%v,CPUs:%v", gi.GoARCH, gi.GoOS, gi.Kernel, gi.Core, gi.Platform, gi.OS, gi.Hostname, gi.CPUs)
}
