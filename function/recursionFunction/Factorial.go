package recursionFunction

type FactorialFunction struct {
}

func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}
