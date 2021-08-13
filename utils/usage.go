package utils

import (
	"edjubert/CodeGenerator/constants"
	"flag"
	"fmt"
	"os"
)

func Usages() {
	fmt.Println("Usages:")
	flag.PrintDefaults()
	os.Exit(constants.WRONG_USAGE)
}
