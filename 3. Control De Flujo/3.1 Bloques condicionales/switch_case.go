package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("La arquitectura de su precesador es ")

	arch := runtime.GOARCH
	switch arch {
	case "386":
		fmt.Println("x85 de 32 bits")
	case "amd64":
		fmt.Println("x86 de 64 bits")
	default:
		fmt.Println(arch)

	}

	fmt.Print("Su sistema operativo es ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("GNU/LINUX")
	case "hurd":
		fmt.Println("GNU/HURD")
	default:
		fmt.Println(os)
	}
}
