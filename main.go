package main

import (
	"fmt"
	"math"
)

func totalAfterInterest(principle float64, rate float64, time float64, frequency float64) float64 {
	return math.Pow(principle*(1+rate/frequency), time*frequency)
}

func promptUser() []float64 {
	var principle float64
	var rate float64
	var time float64
	var frequency float64
	fmt.Println("Please enter the principle amount: ")
	fmt.Scanln(&principle)
	fmt.Println("Please enter the rate: ")
	fmt.Scanln(&rate)
	fmt.Println("Please enter the time: ")
	fmt.Scanln(&time)
	fmt.Println("Please enter the frequency: ")
	fmt.Scanln(&frequency)
	return []float64{principle, rate, time, frequency}
}

func printOutput(total float64, principle float64) {
	fmt.Println("The compound interest is: ", total)
	percentage := total / principle * 100

	fmt.Println("The percentage is: ", percentage)
}

func main() {

	values := promptUser()
	output := totalAfterInterest(values[0], values[1], values[2], values[3])
	printOutput(output, values[0])
}
