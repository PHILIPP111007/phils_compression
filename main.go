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
	numbersCount int = 10_000_000 // 1_000_000_000
)

var (
	workDir                = filepath.Join(CWD, "work_dir")
	fileWithNumbers string = filepath.Join(workDir, "generated_numbers.json")

	subArray types.SubArray
	array    types.ArrayJSON
)

func main() {
	// Create dir
	manager.Manage_dir(workDir)

	manager.Generator(fileWithNumbers, &array, numbersCount)
	subArray = manager.GetRandomArray(100)
	subArray = types.SubArray{'5', '8', '0', '7', '7', '5', '0', '8', '6', '5', '1', '4', '3', '4'}

	manager.Compress(&array, &subArray)
}
