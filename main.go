/*
Главный файл, который следует запускать, в его функции main() содержатся вызовы функций
всех модулей, присутствующих в программе

Проверка работы, ввести число:

834857262215839073214818748531641887316520726390289317717417453874643773795083299576178323786981
*/

package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/db47h/rand64"
	"github.com/db47h/rand64/xorshift"
	"phil.com/manager"
	"phil.com/types"
)

func main() {
	const (
		CWD           string = "/Users/philr/Desktop/гитхаб/Сжатие/Program_Go/"
		numbers_count        = 10_000_000 // 1_000_000_000
		SEED          uint64 = 0          // 1387366483214
	)
	// var file_with_numbers string = filepath.Join(CWD, "generated_numbers", "generated_numbers.txt")

	var source = xorshift.New128plus()
	source.Seed(SEED)
	r64 := rand64.New(source)

	var generated_numbers [][]byte
	for i := 0; i < numbers_count; i++ {
		// generated_numbers = append(generated_numbers, fmt.Sprint(r64.Uint64()))

		generated_numbers = append(generated_numbers, []byte(fmt.Sprint(r64.Uint64())))
	}

	fmt.Println(generated_numbers[:5])

	var generated_numbers_runes []uint8
	for i := 0; i < numbers_count; i++ {

		for _, j := range generated_numbers[i] {

			n, _ := strconv.Atoi(string(j))

			generated_numbers_runes = append(generated_numbers_runes, uint8(n))

		}
	}

	fmt.Println(generated_numbers_runes[:5])

	// file, _ := os.OpenFile(file_with_numbers, os.O_WRONLY, os.ModePerm)
	// file.Write(generated_numbers)
}

func main_1() {
	// ############################################
	// Variables definition
	// ############################################

	const (
		// dir of the `main` file
		CWD string = "/Users/philr/Desktop/гитхаб/Сжатие/Program_Go/"

		// сколько будет сгенерировано чисел,
		// стоит учитывать вычислительные мощности компьютера
		numbers_count int = 100_000_000 // 1_000_000_000
	)

	var (
		dirs = [2]string{
			filepath.Join(CWD, "generated_numbers"),
			filepath.Join(CWD, "work_dir"),
		}
		file_with_numbers string = filepath.Join(dirs[0], "generated_numbers.txt")
		image_RGB         string = filepath.Join(dirs[1], "image_RGB.json")
		generated_numbers types.Channel
	)
	// ############################################
	// ############################################

	// Create dirs
	for _, dir := range dirs {
		manager.Manage_dir(dir)
	}

	generated_numbers = manager.Generator(file_with_numbers, numbers_count)

	// TODO: uncomment
	// Listen user input
	// 	var old_image_path string
	// 	fmt.Print("Введите путь: ")
	// 	_, err := fmt.Scanln(&old_image_path)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// TODO: Заглушка
	old_image_path := filepath.Join(dirs[1], "image.png")

	manager.Manage_file(image_RGB)
	img_data := manager.Get_image_channels(old_image_path, image_RGB)
	manager.Save_image_channels(img_data, image_RGB)
	img_data = manager.Load_image_channels(image_RGB)
	manager.Compress_image_channels(generated_numbers, img_data)
}
