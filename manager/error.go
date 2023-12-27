package manager

func Err_check(err error) {
	if err != nil {
		panic(err)
	}
}
