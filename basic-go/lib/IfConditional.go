package lib

import (
	"fmt"
	"runtime"
)

func If() {
	var os = runtime.GOOS
	if os=="darwin" { 
		fmt.Println("Sistem operasi adalah MacOS")
	} else {
		fmt.Println("Sistem operasi Selain MacOS")
	}
}