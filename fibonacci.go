package fibonacci

func Fibonacci(input int) int {
	if input == 0 {
		return 0
	} else if input == 1 {
		return 1
	} else if input > 1 {
		return Fibonacci(input-1) + Fibonacci(input-2)
	} else {
		return -1
	}		
}