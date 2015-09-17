package main

import ("fmt"; "strconv")
import fibo "Lab1/fibonacci/FibonacciLibrary"

func main() {
	var inputString string
	fmt.Println("Enter the number")	
	fmt.Scanf("%s", &inputString)
	input, err := strconv.Atoi(inputString)
	if err != nil {
		fmt.Println("Invalid input! Please enter a valid integer number!")
	} else {
		output := fibo.Fibonacci(input)
		if output == -1 {
			fmt.Println("Please enter a positive number!")
		} else {
			fmt.Println("Fibonacci sequence number of is : ", output)
		}		
	}	
}