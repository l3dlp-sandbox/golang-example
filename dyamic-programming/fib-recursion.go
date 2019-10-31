package dyamic_programming

func FibRecursion(n int) int {
	if n < 2 {
		return 1
	}
	return FibRecursion(n-1) + FibRecursion(n-2)
}

func FibMemoization(n int) int {
	var memo = make([]int, n)
	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
