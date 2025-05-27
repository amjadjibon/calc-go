package calc

func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / GCD(a, b)
}

func Factorial(n int) int {
	if n < 0 {
		panic("factorial function requires a non-negative integer")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
