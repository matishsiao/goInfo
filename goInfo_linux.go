package goInfo

import (
	"runtime"
)

func GetInfo() (GoInfoObject, error) {
	kernel, core, platform, osname, err := _getInfo("-srio")
	gio := GoInfoObject{
		Kernel:   kernel,
		Core:     core,
		Platform: platform,
		OS:       osname,
		GoOS:     runtime.GOOS,
		CPUs:     runtime.NumCPU(),
		Hostname: _hostname(),
	}
	return gio, err
}
