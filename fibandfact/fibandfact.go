package fibandfact

func Fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return Fib(n-2) + Fib(n-1)
	}
}

func Fact(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}
