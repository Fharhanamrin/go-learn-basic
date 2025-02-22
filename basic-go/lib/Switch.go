package lib

import (
	"runtime"
	"fmt"
)


func Switch() {		
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Sistem operasi adalah MacOS")
	case "linux":
		fmt.Println("Sistem operasi adalah Linux")
	default:
		fmt.Println("Sistem operasi adalah", os)
	}
}