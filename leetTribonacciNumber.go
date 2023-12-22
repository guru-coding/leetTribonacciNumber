package leetTribonacciNumber

func tribonacci(n int) int {
	dpc := make([]int, 3)
	dpc[0] = 0
	dpc[1] = 1
	dpc[2] = 1
	for i := 3; i <= n; i++ {
		dpc[i%3] = dpc[(i-1)%3] + dpc[(i-2)%3] + dpc[(i-3)%3]
	}
	return dpc[n%3]
}
