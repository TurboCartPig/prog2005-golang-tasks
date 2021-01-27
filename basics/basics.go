package basics

// Task 0.
// ********************************************************************
func Swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// Task 1.
// ********************************************************************
func Fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return Fib(n-2) + Fib(n-1)
	}
}

// Task 2.
// ********************************************************************
func Fact(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}

// Task 5.
// ********************************************************************
func GiveMeFunc() func() func(a, b int) int {
	return func() func(a, b int) int { return func(a, b int) int { return a + b } }
}

// Task 6.
// ********************************************************************
type Person struct {
	age int
}

// Takes a copy of a person and returns that copy modified
func grow1(p Person) Person {
	p.age += 10
	return p
}

// Takes a pointer to a person and modifies through that pointer
func grow2(p *Person) {
	p.age += 10
}

// Receives a copy of a person and returns that copy modified
func (p Person) grow3() Person {
	p.age += 10
	return p
}

// Recives a pointer to a person and modifies through that pointer
func (p *Person) grow4() {
	p.age += 10
}
