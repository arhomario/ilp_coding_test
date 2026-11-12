package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("===SOAL 1===")
	// fmt.Println(a([]int{-1, 1}))
	// fmt.Println(a([]int{-1, -7, -5}))
	// fmt.Println(a([]int{1, 2, 1, 6}))

	fmt.Println("===SOAL 2===")
	fmt.Println(b([]int{0, 1, 0}, 1000, 1010))
	fmt.Println(b([]int{0, 1, 0}, 1259, 2601))

}

func a(arr []int) int {

	hasNegative := slices.ContainsFunc(arr, func(n int) bool {
		return n < 0
	})

	hasPositive := slices.ContainsFunc(arr, func(n int) bool {
		return n > 0
	})

	var result int
	if hasNegative && hasPositive {
		result = 0
	} else if hasNegative {
		result -= len(arr) - 1
	} else if hasPositive {
		result = len(arr) - 1
	}

	return result
}

func b(exclude []int, nums ...int) (int, error) {

	fmt.Println(strings.Join(strconv.Itoa(nums), ","))
	for _, nilai := range nums {
		fmt.Println(nilai)
	}

	for i := 0; i < len(exclude); i++ {
		fmt.Println(exclude[i])
	}

	return 100, fmt.Errorf("not found")
}
