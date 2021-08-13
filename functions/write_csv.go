package functions

import (
	"bufio"
	"os"
	"strconv"

	"edjubert/CodeGenerator/utils"
)

const ID_LEN = 12

func WriteCSV(filename string, nLines int, prefix string, ext string) {
	dir, err := os.Getwd()
	utils.Check(err)

	extension := ext
	if ext != "" {
		extension = "." + ext
	}
	file := dir + "/" + filename + extension
	f, err := os.Create(file)
	utils.Check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	for i := 1; i <= nLines; i++ {
		str := RandStringRunes(ID_LEN)
		_, err := w.WriteString(prefix + "_" + str + "_" + strconv.Itoa(i) + "\n")
		utils.Check(err)

		if i%10 == 0 {
			w.Flush()
		}
	}
	w.Flush()

}
