package fibonacci

import "testing"

type testFibonacci struct {
	input int
	expectedOutput int
}

var tests = [] testFibonacci {
	{3, 2},
	{4, 3},
	{5, 5},
	{9, 34},
}

func TestFibonacci(t *testing.T) {
	for _, test := range tests {
		actualResult := Fibonacci(test.input) 
		if actualResult != test.expectedOutput {
			t.Error("Fibonacci (%d) : Actual Result - (%d), Expected Result - (%d)", test.input, actualResult, test.expectedOutput)
		}
	}
}	
