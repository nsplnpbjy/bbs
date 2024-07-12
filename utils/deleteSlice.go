package utils

func DeleteSlice[T string | int | int64](slice []T, ele T) []T {
	for i := 0; i < len(slice); i++ {
		if slice[i] == ele {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}
