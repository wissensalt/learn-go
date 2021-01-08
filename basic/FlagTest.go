package main

import (
	"flag"
	"fmt"
)

func main() {
	/**
		First Param : flag name
		Second Param : default value
		Third Param : description/ help text
	 */
	archPtr := flag.String("arch", "x86", "CPU Type")
	flag.Parse()

	switch *archPtr {
		case "x86":
			fmt.Println("Running in 32 bit mode")
		case "AMD64":
			fmt.Println("Running in 64 bit mode")
		case "IA64":
			fmt.Println("Running in IA64")
	}

	fmt.Println("Process Completed")
}
