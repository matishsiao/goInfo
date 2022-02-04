package goInfo

import (
	"runtime"
)

func GetInfo() (GoInfoObject, error) {
	kernel, core, platform, _, err := _getInfo("-srm")
	gio := GoInfoObject{
		Kernel:   kernel,
		Core:     core,
		Platform: platform,
		OS:       kernel,
		GoOS:     runtime.GOOS,
		CPUs:     runtime.NumCPU(),
		Hostname: _hostname(),
	}
	return gio, err
}
