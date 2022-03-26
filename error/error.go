package error

func PanicErr[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

// func LogErr()
