package array

func Has[T comparable](arr []T, target T) bool {
	for _, it := range arr {
		if it == target {
			return true
		}
	}
	return false
}
