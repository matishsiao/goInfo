package goInfo

import (
	"fmt"
)

type GoInfoObject struct {
	Kernel string
	Core string
	Platform string
	OS string
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("Kernel:",gi.Kernel)
	fmt.Println("Core:",gi.Core)
	fmt.Println("Platform:",gi.Platform)
	fmt.Println("OS:",gi.OS)
}

