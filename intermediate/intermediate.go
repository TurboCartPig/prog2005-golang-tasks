package intermediate

import "fmt"

// Task 1
// Guessing the task wants me to use a switch or if/else...
func Hex(i string) string {
	var o string

	for _, c := range i {
		var suffix string
		switch c {
		case 'a':
			suffix = "10"
		case 'b':
			suffix = "11"
		case 'c':
			suffix = "12"
		case 'd':
			suffix = "13"
		case 'e':
			suffix = "14"
		case 'f':
			suffix = "15"
		}
		o += suffix
	}

	return o
}

// Task 2
// This is probably a bit silly
func Hex2(i string) string {
	var o string

	for _, c := range i {
		var i int
		fmt.Sscanf(string(c), "%x", &i)
		o += fmt.Sprint(i)
	}

	return o
}

// Task 4
type Person struct {
	age int
}

func (self *Person) Apply(f func(*Person)) {
	f(self)
}

// Task 5

// Task 6?
// Given a string composed of a sequence of integers, e.g. "1 4 5 2 4 3 1 2 1" find out which integer
// occurs the most number of times, and print it out. If there are two that have the same count,
// print the smaller one. For the above example, the answer is 1.
