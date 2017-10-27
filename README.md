# Introduction:
  GoInfo is get os platform information coding by Golang.
  
  It can help you to know os information.
  
  
## Version:

version:0.0.1

## Futures

get linux information

get windows information

get osx information

get freebsd information
  
  
## Install:
```sh
  go get github.com/matishsiao/goInfo
  go build
```

## Struct:
```sh
  type GoInfoObject struct {
	GoOS string
	Kernel string
	Core string
	Platform string
	OS string
	Hostname string
	CPUs int
}
```

## Example:

```sh   
   package main

   import (
	   "github.com/matishsiao/goInfo"
   )

   func main() {
		gi := goInfo.GetInfo()
		gi.VarDump()
	 }

```
Get Linux Distribution version
```sh
package main
import (
     "github.com/pyToshka/goInfo"
)

func main() {
    dist :=goInfo.GetInfo().Distribution
    print(dist)


}
```
Return
```sh
ubuntu@ubuntu-xenial:/tmp$
16.04
```

Get Linux Distribution name
```sh
package main
import (
     "github.com/pyToshka/goInfo"
)

func main() {
	dist :=goInfo.GetInfo().Name
	print(dist)


}
```
Return
```sh
[vagrant@localhost tmp]$
centos
```

It's will show:

```sh
   GoOS: linux
   Kernel: Linux
   Core: 3.10.0-693.2.1.el7.x86_64
   Platform: x86_64
   OS: GNU/Linux
   Hostname: localhost.localdomain
   CPUs: 1
   Distribution version: 7
   Distribution name: centos
```

##License and Copyright
This software is Copyright 2012-2014 Matis Hsiao.
