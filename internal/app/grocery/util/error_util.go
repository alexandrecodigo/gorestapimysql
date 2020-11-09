package util

// PanicError error
func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
