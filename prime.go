package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var nums int
	fmt.Println("How many numbers would you like to find?")
	_, _ = fmt.Scan(&nums)

	fmt.Println("Enter a number.")
	arr := make([]int, nums)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	getIntSlice(input, &arr, nums)

	for i := 0; i < len(arr); i++ {
		isPrime(arr[i])
	}

	// n, _ := fmt.Scanln(&num)
	// for i:=0; i<n;i++{
	//
	// }

}
func getIntSlice(input string, arr *[]int, n int) {
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		*arr = append(*arr, num)
	}
}

// func inputInt(n int) ([]int, error) {

// 	x := make([]int, n)

// 	for i := range x {
// 		_, err := fmt.Scan(&x[i])
// 		if err != nil {
// 			fmt.Println(err)
// 			return x[:i], err
// 		}
// 	}
// 	return x, nil
// }

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

				fmt.Println("It isn't a Prime number.")
				return
			}
		}
	}

	fmt.Println("It is a Prime number.")

}
