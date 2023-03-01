package perf

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func tailFact(n int) int {
	return factT(n-1, n)
}

func factT(n, current int) int {
	if n == 1 {
		return current
	}
	return factT(n-1, n*current)
}
