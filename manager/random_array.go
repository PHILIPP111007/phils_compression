package manager

import (
	"fmt"
	"math/rand"

	"phil.com/types"
)

func GetRandomArray(length int) (subArray types.SubArray) {

	for i := 0; i < length; i++ {

		n := rand.Intn(10)
		s := fmt.Sprint(n)
		b := []byte(s)

		subArray = append(subArray, b[0])
	}

	return subArray
}
