package h

// Poe 如果有错误，则 panic 并终止程序。
// If there is an error, panic and terminate the program.
func Poe(err error) {
	if err != nil {
		panic(err)
	}
}
