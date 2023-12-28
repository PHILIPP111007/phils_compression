package manager

import (
	"fmt"

	"phil.com/types"
)

var (
	ArrayLen, subArrayLen int
	ArrayLenTable         types.ArrayLenTable
)

func Compress(array types.Array, subArray types.SubArray) {

	setSubArrayLen(subArray)
	setArrayLenAndArrayLenTable(array)

	for {
		if subArrayLen == 0 {
			break
		}

		n, x, y := recursion(array, subArray, 1, 0, 0)
		fmt.Println(subArrayLen, n, x, y)

		subArray = subArray[n:]
		subArrayLen -= n
	}
}

func recursion(array types.Array, subArray types.SubArray, n, x, y int) (int, int, int) {
	var x_new, y_new = getIndex(array, subArray[:n], n)

	if x_new != -1 && n < subArrayLen {
		return recursion(array, subArray, n+1, x_new, y_new)
	}
	return n, x, y
}

func getIndex(array types.Array, subArray types.SubArray, subArrayLen int) (int, int) {
	return Index(array, subArray, subArrayLen)
}

func setArrayLenAndArrayLenTable(array types.Array) {
	for _, elem := range array.Value {
		ArrayLen += len(elem)
		ArrayLenTable = append(ArrayLenTable, len(elem))
	}
}

func setSubArrayLen(subArray types.SubArray) {
	subArrayLen = len(subArray)
}
