package perr

func PanicErrDouble[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
