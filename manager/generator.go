/*
Данный модуль генерирует максимально длинную последовательность
чисел, длина зависит от вычеслительных мощностей компьютера.
*/

package manager

import (
	"encoding/json"
	"fmt"

	"phil.com/types"

	"io"
	"os"

	rand64 "github.com/db47h/rand64"
	xorshift "github.com/db47h/rand64/xorshift"
)

const SEED = 0 // 1387366483214

func Generator(file_with_numbers string, array *types.Array, numbers_count int) {

	if is_created_before := Manage_file(file_with_numbers); !is_created_before {
		generate(array, numbers_count)

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

		err = json.Unmarshal(byteArray, &array)
		Err_check(err)

		err = file.Close()
		Err_check(err)
	}
}

func generate(array *types.Array, numbers_count int) {
	var source = xorshift.New128plus()
	source.Seed(SEED)
	r64 := rand64.New(source)

	for i := 0; i < numbers_count; i++ {
		array.Value = append(array.Value, []byte(fmt.Sprint(r64.Uint64())))
	}
}
