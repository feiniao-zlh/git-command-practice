package utils

func Add(a, b int) int {
	return a + b
}

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
