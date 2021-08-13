package main

import (
	eParse "edjubert/CodeGenerator/parse"
	"fmt"
)

func main() {
	n, name, prefix, ext := eParse.Flags()
	fmt.Println(n, name, prefix, ext)
}
