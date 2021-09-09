package main

import "fmt"

func main() {
	FibonacciNumber := 15

	fmt.Println("Fibonnaci's numbers, first " + fmt.Sprint(FibonacciNumber) +
		" numbers in the sequence: ")
	printValsInSlice(fibonacciSlice(FibonacciNumber))
}

func fibonacciSlice(n int) []int {

	fibSlice := []int{}

	switch x := n; {

	case x >= 1:
		fibSlice = append(fibSlice, 0)
		fallthrough

	default:
		for i := 2; i <= n; i++ {
			fibSlice = append(fibSlice, fibonacciVal(i))
		}
	}

	return fibSlice
}

func fibonacciVal(fibonacciNum int) int {
	if fibonacciNum == 2 {
		return 1
	} else if fibonacciNum == 1 {
		return 0
	}
	return fibonacciVal(fibonacciNum-1) + fibonacciVal(fibonacciNum-2)
}

func printValsInSlice(sliceVals []int) {
	for i := 0; i < len(sliceVals); i++ {
		fmt.Print(sliceVals[i])
		if i >= 0 && i < len(sliceVals)-1 {
			fmt.Print(", ")
		}
	}
}
