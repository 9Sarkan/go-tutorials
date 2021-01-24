package benchmark

func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}
func fib3(n int) int {
	result := make(map[int]int)
	for i := 0; i <= n; i++ {
		if i <= 2 {
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result[n]
}
