package main

import ("fmt"; "time")
import sleep "Lab1/sleepfunction/SleepLibrary"

func main() {
	var duration time.Duration
	fmt.Println("Enter time duration to show sum of 10 and 5")
	fmt.Scanf("%d", &duration)
	if(duration > 0) {
		sleep.Sleep(duration)
	} else {
		fmt.Println("Please enter valid time!")
	}
}