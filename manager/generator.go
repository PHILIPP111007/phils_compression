/*
Данный модуль генерирует максимально длинную последовательность
чисел, длина зависит от вычеслительных мощностей компьютера.
*/

package manager

import (
	"encoding/json"
	"fmt"
	"io"

	"phil.com/types"

	"os"

	rand64 "github.com/db47h/rand64"
	xorshift "github.com/db47h/rand64/xorshift"
)

const SEED = 0 // 1387366483214

func Generator(file_with_numbers string, array *types.ArrayJSON, numbersCount int) {
	if is_created_before := Manage_file(file_with_numbers); !is_created_before {
		generate(array, numbersCount)

		file, err := os.OpenFile(file_with_numbers, os.O_WRONLY, os.ModePerm)
		Err_check(err)

		byteArray, err := json.Marshal(array)
		Err_check(err)

		_, err = file.Write(byteArray)

		Err_check(err)
	} else {
		file, err := os.OpenFile(file_with_numbers, os.O_RDONLY, os.ModePerm)
		Err_check(err)

		byteArray, err := io.ReadAll(file)
		Err_check(err)

		err = json.Unmarshal(byteArray, array)
		Err_check(err)

		err = file.Close()
		Err_check(err)
	}
}

func generate(array *types.ArrayJSON, numbersCount int) {
	var source = xorshift.New128plus()
	source.Seed(SEED)
	r64 := rand64.New(source)

	var array_1 = make(types.Array, numbersCount)
	for i := 0; i < numbersCount; i++ {

		num := r64.Uint64()

		go func(i int, num uint64) {
			array_1[i] = []byte(fmt.Sprint(num))
		}(i, num)
	}

	array.Value = array_1
}
