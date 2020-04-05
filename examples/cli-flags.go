package main

import (
	"flag"
	"fmt"
)

func main() {

	strPtr := flag.String("default-user", "guest", "default username")
	numPtr := flag.Int("log-level", 2, "log level")
	var boolPtr bool
	flag.BoolVar(&boolPtr, "disable-log", false, "disable logs")

	flag.Parse()

	fmt.Println(*strPtr)
	fmt.Println(*numPtr)
	fmt.Println(*&boolPtr)
}
