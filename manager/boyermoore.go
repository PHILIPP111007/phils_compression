package manager

import (
	"fmt"

	"phil.com/types"
)

// Index returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func Index(array types.Array, subArray types.SubArray, subArrayLen int) (int, int) {
	d := calculateSlideTable(subArray, subArrayLen)
	return indexWithTable(&d, array, subArray, subArrayLen)
}

// IndexWithTable returns the first index substr found in the s.
// It needs the slide information of substr
func indexWithTable(d *[256]int, array types.Array, subArray types.SubArray, subArrayLen int) (int, int) {
	switch {
	case subArrayLen == 0:
		return 0, 0
	case subArrayLen > ArrayLen:
		return -1, -1
	case subArrayLen == ArrayLen:

		// not optimized

		var arrayUnited string
		var subArrayUnited string

		for _, elem := range subArray {
			subArrayUnited += fmt.Sprint(elem)
		}

		for _, arr := range array.Value {
			for _, elem := range arr {
				arrayUnited += fmt.Sprint(elem)
			}
		}

		if arrayUnited == subArrayUnited {
			return 0, 0
		}
		return -1, -1
	}

	i := 0
	J := subArrayLen - 1
	for i <= ArrayLen-subArrayLen {
		j := J

		// Ищем вхождение
		x, y := getArrayIndexes(i + j)
		for j >= 0 && subArray[j] == array.Value[x][y] {
			j--

			if j < 0 {
				return x, y // Найдено вхождение
			}

			x, y = getArrayIndexes(i + j)
		}

		// Смещение в соответствии с таблицей плохих символов
		slid := j - d[array.Value[x][y]]
		if slid < 1 {
			slid = 1
		}
		i += slid
	}

	return -1, -1 // Вхождение не найдено
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func calculateSlideTable(subArray types.SubArray, subArrayLen int) [256]int {

	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < subArrayLen; i++ {
		d[subArray[i]] = i
	}
	return d
}

func getArrayIndexes(sum_i_j int) (x int, y int) {

	for x, arrLen := range ArrayLenTable {

		sum_i_j -= arrLen

		if sum_i_j < 0 {
			return x, (-1 * sum_i_j) - 1
		}
	}

	return 0, 0
}
