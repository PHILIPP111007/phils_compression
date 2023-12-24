/*
Данный модуль генерирует максимально длинную последовательность
чисел, длина зависит от вычеслительных мощностей компьютера.
*/

package manager

// go get github.com/db47h/rand64

import (
	"fmt"
	"log"
	"os"
	"strings"

	rand64 "github.com/db47h/rand64"
	xorshift "github.com/db47h/rand64/xorshift"
)

const SEED uint64 = 1387366483214

func generate_numbers(file_path string, insert_spaces bool, first_range, second_range int) {
	var source = xorshift.New128plus()
	source.Seed(SEED)
	r64 := rand64.New(source)

	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := 1; i <= first_range; i++ {
		fmt.Printf("%d / %d\n", i, first_range)

		lineArr := make([]string, second_range)
		for j := 0; j < second_range; j++ {
			lineArr[j] = fmt.Sprint(r64.Uint64())
		}

		var line string
		if insert_spaces {
			line = strings.Join(lineArr, " ")
		} else {
			line = strings.Join(lineArr, "")
		}
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func create_file(file_path string) (is_created_before bool) {
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		file, err := os.Create(file_path)
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()
		return false
	} else {
		return true
	}
}

func Generator(files_from_generator []string, new_numbers_count, first_range int) {
	var second_range int = new_numbers_count / first_range

	for i, file_path := range files_from_generator {
		is_created_before := create_file(file_path)
		if !is_created_before {
			fmt.Printf("\nStarted creating: %s\n", file_path)

			if i == 0 {
				generate_numbers(file_path, false, first_range, second_range)
			} else {
				generate_numbers(file_path, true, first_range, second_range)
			}
		}
	}

	fmt.Println("Generator: Done.")
}
