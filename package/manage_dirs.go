package manager

import (
	"fmt"
	"os"
)

func Manage_dirs() {
	// path/to/whatever does not exist
	if _, err := os.Stat("generated_numbers/"); os.IsNotExist(err) {
		os.Mkdir("generated_numbers", os.ModePerm)
	}

	if _, err := os.Stat("work_dir/"); os.IsNotExist(err) {
		os.Mkdir("work_dir", os.ModePerm)
	}

	fmt.Println("Manage dirs: Done.")
}
