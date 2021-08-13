package parse

import (
	"errors"
	"flag"
	"fmt"

	"edjubert/CodeGenerator/utils"
)

func Flags() (*int, *string, *string, *string) {
	nLines := flag.Int("nLines", 0, "Number of line you want to create for your CSV")
	fileName := flag.String("name", "", "Your filename")
	prefix := flag.String("prefix", "", "The prefix used to generate the CSV with format <prefix>_<line_number>")
	ext := flag.String("ext", "", "The file extension you want to create")

	flag.Parse()

	if *nLines <= 0 || *fileName == "" || *prefix == "" {
		fmt.Println(errors.New("ERROR: Missing args\n"))
		utils.Usages()
	}

	return nLines, fileName, prefix, ext

}
