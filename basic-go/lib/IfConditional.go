package lib

import (
	"fmt"
	"runtime"
)

func If() {
	var os = runtime.GOOS
	if os=="darwin" { 
		fmt.Println("Sistem os MacOS")
	} else {
		fmt.Println("Sistem os selain MacOS")
	}
}