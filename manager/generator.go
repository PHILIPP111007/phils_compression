/*
Данный модуль генерирует максимально длинную последовательность
чисел, длина зависит от вычеслительных мощностей компьютера.
*/

package manager

// go get github.com/db47h/rand64

import (
	"fmt"

	"phil.com/types"

	"io"
	"os"

	rand64 "github.com/db47h/rand64"
	xorshift "github.com/db47h/rand64/xorshift"
)

const SEED uint64 = 1387366483214

func Generator(file_with_numbers string, numbers_count int) types.Channel {
	var generated_numbers types.Channel

	if is_created_before := Manage_file(file_with_numbers); !is_created_before {
		generated_numbers = generate(numbers_count)
		file, err := os.OpenFile(file_with_numbers, os.O_WRONLY, os.ModePerm)
		Err_check(err)

		_, err = file.Write(generated_numbers)
		Err_check(err)
	} else {
		file, err := os.OpenFile(file_with_numbers, os.O_RDONLY, os.ModePerm)
		Err_check(err)

		generated_numbers, err = io.ReadAll(file)
		Err_check(err)

		err = file.Close()
		Err_check(err)
	}

	return generated_numbers
}

func generate(numbers_count int) types.Channel {
	var source = xorshift.New128plus()
	source.Seed(SEED)
	r64 := rand64.New(source)

	var generated_numbers types.Channel

	for i := 0; i < numbers_count; i++ {
		generated_numbers = append(generated_numbers, []byte(fmt.Sprint(r64.Uint64()))...)
		// for _, j := range fmt.Sprint(r64.Uint64()) {
		// 	generated_numbers = append(generated_numbers, byte(j))
		// }
	}

	return generated_numbers
}
