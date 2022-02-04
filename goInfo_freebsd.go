package goInfo

import (
	"runtime"
)

func GetInfo() (GoInfoObject, error) {
	kernel, core, platform, _, err := _getInfo("-sri")
	gio := GoInfoObject{
		Kernel:   kernel,
		Core:     core,
		Platform: runtime.GOARCH,
		OS:       platform,
		GoOS:     runtime.GOOS,
		CPUs:     runtime.NumCPU(),
		Hostname: _hostname(),
	}
	return gio, err
}
