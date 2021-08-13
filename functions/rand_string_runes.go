package functions

import (
	"edjubert/CodeGenerator/constants"
	"math/rand"
)

// func InitRand() {
// 	rand.Seed(time.Now().UnixNano())
// }

var letterRunes = []rune(constants.LetterRunes)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}
