package main

import (
	"edjubert/CodeGenerator/functions"
	"edjubert/CodeGenerator/parse"
)

func main() {
	n, name, prefix, ext := parse.Flags()

	functions.WriteCSV(*name, *n, *prefix, *ext)
}
