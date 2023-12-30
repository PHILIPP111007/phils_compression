package manager

import (
	"os"
)

func Manage_file(file_path string) (is_created_before bool) {
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		file, err := os.Create(file_path)
		Err_check(err)

		err = file.Close()
		Err_check(err)

		return false
	}
	return true
}

func Manage_dir(dir string) {
	// path/to/whatever does not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
}
