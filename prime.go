package main

import (
	"fmt"
	"math"
)

func main() {

	var nums int
	fmt.Println("How many numbers would you like to find?")
	_, _ = fmt.Scan(&nums)

	var numbers []int
	fmt.Println("Enter a number.")
	numbers, _ = inputInt(nums)

	for i := 0; i < len(numbers); i++ {
		isPrime(numbers[i])
	}

	// n, _ := fmt.Scanln(&num)
	// for i:=0; i<n;i++{
	//
	// }

}

func inputInt(n int) ([]int, error) {

	x := make([]int, n)

	for i := range x {
		_, err := fmt.Scan(&x[i])
		if err != nil {
			fmt.Println(err)
			return x[:i], err
		}
	}
	return x, nil
}

func isPrime(x int) {

	if x < 2 {
		fmt.Println("It isn't a Prime number.")
		return
	} else if x == 2 {

		fmt.Println("It is a Prime number.")
		return
	} else {

		max := int(math.Sqrt(float64(x)))
		for i := 2; i <= max; i++ {

			if x%i == 0 {

				fmt.Print("It isn't a Prime number.")
				return
			}
		}
	}

	fmt.Print("It is a Prime number.")

}
