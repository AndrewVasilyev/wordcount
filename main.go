package main

import (
	"flag"
	"fmt"

	//"fmt"
	"strings"
)

func main() {
	flag.Parse()
	src := strings.Join(flag.Args(), "")

	samples := strings.Split(src, " ")

	if samples[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(samples))
	}
}
