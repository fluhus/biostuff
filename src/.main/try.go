package main

import (
	"fmt"
	"strings"
	"os"
)

func pe(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	pe("start")
	
	s := "amit\nlavon"
	pe(strings.Split(s, ""))
	
	pe("end")
}










