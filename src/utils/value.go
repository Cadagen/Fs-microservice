package utils

func Ifs[T any](cond bool, truev, falsev T) T {
	if cond {
		return truev
	}

	return falsev
}
