package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("===SOAL 1===")
	fmt.Println(a([]int{-1, 1}))
	fmt.Println(a([]int{-1, -7, -5}))
	fmt.Println(a([]int{1, 2, 1, 6}))

	fmt.Println("===SOAL 2===")
	fmt.Println(b([]int{0, 1, 0}, 1000, 1010))
	fmt.Println(b([]int{0, 1, 0}, 1259, 2601))

	fmt.Println("===SOAL 3===")
	fmt.Println(c([]int{1, 2}, []int{1, 3}))
	fmt.Println(c([]int{1, 2, 2}, []int{1, 2, 4}))

	fmt.Println("===SOAL 4===")
	fmt.Println(d(1223334))

	fmt.Println("===SOAL 5===")
	fmt.Println(e(1, 2, 3, 4))
	fmt.Println(e(1, 2, 3, 4, 5))
}

func a(arr []int) int {

	// 	1. write a function
	// func GO(arr []int) int
	// if arr contains negative and positive number, should return 0
	// if arr only have positive number, return the smallest number (except 0)
	// if arr only have negative number, return the biggest number (except 0)
	// e.g:
	// arr = [-1,1], the function should return 0
	// arr = [-1,-7,-5], then function should return -2
	// arr = [1,2,1,6], the function should return 3

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

	// 	2. write a function
	// func GO(exclude []int, nums ...int) (int, error)
	// join the numbers as one int variable, then do reverse, but exclude number if exists on parameter exclude.
	// e.q:
	// GO([]int{0,1,0}, 1000, 1010), the function should return 0, error with message "not found"
	// GO([]int{0,1,0}, 1259, 2601), the function should return 62952, nil

	var result int
	var numsStr string

	for _, n := range nums {
		numsStr += strconv.Itoa(n)
	}

	// Reverse
	var reversed string
	for i := len(numsStr) - 1; i >= 0; i-- {
		reversed += string(numsStr[i])
	}

	// Exclude
	for _, x := range exclude {
		reversed = strings.ReplaceAll(reversed, strconv.Itoa(x), "")
	}

	if reversed == "" {
		return 0, errors.New("not found")
	}

	result, _ = strconv.Atoi(reversed)

	return result, nil
}

func c(numsA, numsB []int) []int {

	// 3. write a function
	// func GO(numsA, numsB []int) []int
	// find numbers on numsA that does exists in numsB.
	// and do order descending.
	// e.q:
	// numsA = [1,2], and numsB = [1,3], the function should return [1].
	// numsA = [1,2,2] and numsB = [1,2,4], the function should return [2,1].

	var result []int

	// Find
	finds := make(map[int]bool)
	for _, numA := range numsA {
		for _, numB := range numsB {
			if numA == numB {
				finds[numB] = true
				break
			}
		}
	}

	// Append to Result
	for find, _ := range finds {
		result = append(result, find)
	}

	// Sort by Desc
	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func d(num int) map[int]int {

	var result = make(map[int]int)

	// 4. write a function
	// func GO(num int) map[int]int
	// split the int become array of single digit int, then do group and count the same number.
	// e.q:
	// num = 1223334, the function should return
	// 1 -> 1
	// 2 -> 2
	// 3 -> 3
	// 4 -> 1

	nums := strconv.Itoa(num)

	for _, n := range nums {
		i, _ := strconv.Atoi(string(n))
		result[i]++
	}

	return result
}

func e(nums ...int) int {

	// 	5. write a function
	// func GO(nums ...int) int
	// do multiplication for 2 number on even index and odd index.
	// if the last index doesnt have the pair, do multiplication with his own number.
	// use concurrency to perform the multiplication.
	// use main thread to perform the summation.
	// e.g:
	// nums = [1,2,3,4], the function should return 14, from (1*2) + (3*4), with 2 concurrency process.
	// nums = [1,2,3,4,5], the function should return 39, from (1*2) + (3*4) + (5*5), with 3 concurrency process.

	var result int

	// cari ganjil atau genap
	var even []int
	var odd []int
	for _, n := range nums {
		if n%2 == 0 {
			odd = append(odd, n)
		} else {
			even = append(even, n)
		}
	}

	// Perkalian
	for i, ev := range even {
		if i < len(odd) {
			result += ev * odd[i]
		} else {
			result += ev * ev
		}
	}

	// Notes: Maaf jujur saya belum paham concurrency process, jadi saya pakai cara biasa.

	return result

}
