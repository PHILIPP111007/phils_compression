/*
Главный файл, который следует запускать, в его функции main() содержатся вызовы функций
всех модулей, присутствующих в программе

Проверка работы, ввести число:

subArray = types.SubArray{'5', '8', '0', '7', '7', '5', '0', '8', '6', '5', '1', '4', '3', '4'}
*/

package main

import (
	"path/filepath"

	"phil.com/manager"
	"phil.com/types"
)

const (
	// dir of the `main` file
	CWD string = "/Users/philr/Desktop/гитхаб/Сжатие/Program_Go/"

	// сколько будет сгенерировано чисел,
	// стоит учитывать вычислительные мощности компьютера
	numbersCount int = 10_000 // 1_000_000_000
)

var (
	dirs = [2]string{
		filepath.Join(CWD, "generated_numbers"),
		filepath.Join(CWD, "work_dir"),
	}
	fileWithNumbers string = filepath.Join(dirs[0], "generated_numbers.txt")

	subArray types.SubArray
	array    types.Array
)

func main() {
	// Create dirs
	for _, dir := range dirs {
		manager.Manage_dir(dir)
	}

	manager.Generator(fileWithNumbers, &array, numbersCount)
	subArray = manager.GetRandomArray(100)
	manager.Compress(array, subArray)
}
