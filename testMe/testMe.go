package testme

func f1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return f1(n-1) + f1(n-2)
}
func f2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	return f2(n-1) + f2(n-2)
}

func s1(str string) int {
	if str == "" {
		return 0
	}
	var len = 1
	for range str {
		len++
	}
	return len
}
func s2(str string) int { return len(str) }
