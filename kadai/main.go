package main

import "fmt"

func main() {
	oneToTen := newInts()

	for _, num := range oneToTen {
		if num%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}

	}

}

func newInts() []int {
	retVal := []int{}
	for i := 0; i <= 10; i++ {
		retVal = append(retVal, i)
	}

	return retVal
}
