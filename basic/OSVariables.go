package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("We are running GO using version %v within OS %v\n", runtime.Version(), runtime.GOOS)
}
