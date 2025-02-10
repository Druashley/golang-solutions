package sevenkyu

func Factorial(n int) int {
	answer := 1
	if n == 0 {
		answer = 1
	}

	for i := 1; i <= n; i++ {
		answer = i * answer
	}

	return answer
}
