/*
Главный файл, который следует запускать, в его функции main() содержатся вызовы функций
всех модулей, присутствующих в программе

Проверка работы, ввести число:

834857262215839073214818748531641887316520726390289317717417453874643773795083299576178323786981
*/

package main

import (
	manager "phil.com/package"
)

var files_from_generator = []string{"generated_numbers/numbers_from_generator_1.txt", "generated_numbers/numbers_from_generator_2.txt"}

const new_numbers_count = 1_000_000 // 100_000_000 // сколько будет сгенерировано чисел, можно изменить но стоит учитывать вычислительные мощности компьютера
const first_range = 10              // 10_000 // можно изменить, когда переменная new_numbers_count слишком велика, для нормальной нагрузки на память
const file_RGB = "work_dir/RGB.json"

func main() {

	// manager.Manage_dirs()
	// manager.Generator(
	// 	files_from_generator,
	// 	new_numbers_count,
	// 	first_range,
	// )

	// fmt.Print("обработать число (1)\nобработать картинку (2)\n: ")
	// var choice string

	// _, err := fmt.Scanln(&choice)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// switch choice {

	// case "1":
	// 	var number string
	// 	fmt.Print("Введите число: ")
	// 	_, err := fmt.Scanln(&number)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	manager.Work_from_number(number, file_RGB)

	// case "2":
	// 	var image_path string
	// 	fmt.Print("Введите путь: ")
	// 	_, err := fmt.Scanln(&image_path)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	manager.Work_from_image(image_path, file_RGB)
	// }

	image_path := "work_dir/image.png"
	manager.Work_from_image(image_path, file_RGB)
	// fmt.Println("Main: Done.")
}
