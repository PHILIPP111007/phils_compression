package manager

import (
	"fmt"
	"math/rand"

	"phil.com/types"
)

func GetRandomArray(length int) types.SubArray {

	var arr types.SubArray

	for i := 0; i < length; i++ {

		n := rand.Intn(10)
		s := fmt.Sprint(n)
		b := []byte(s)

		arr = append(arr, b[0])
	}

	return arr
}
