package lib

import (
	"runtime"
	"fmt"
)


func Switch() {		
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Sistem operasi MacOS")
	case "linux":
		fmt.Println("Sistem operasi Linux")
	default:
		fmt.Println("Sistem operasi", os)
	}
}