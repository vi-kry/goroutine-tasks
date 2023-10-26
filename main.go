package main

import (
	"fmt"
	"sync"
)

// 1.Параллельная обработка слайса чисел:
// Есть слайс чисел, и нужно параллельно вычислить сумму, минимум и максимум этого слайса.
// Т.е три горутины которые используют один слайс

func main() {
	nums := []int{2, 1, 5, 4, 3}

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		findSum(nums)
	}()

	go func() {
		defer wg.Done()
		findMin(nums)
	}()

	go func() {
		defer wg.Done()
		findMax(nums)
	}()

	wg.Wait()

}

func findSum(nums []int) {

	var sum int

	for _, n := range nums {
		sum += n
	}
	fmt.Printf("sum = %d\n", sum)
}

func findMin(nums []int) {

	min := nums[0]

	if len(nums) == 0 {
		min = -1
	}

	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	fmt.Printf("min = %d\n", min)
}

func findMax(nums []int) {

	max := nums[0]

	if len(nums) == 0 {
		max = -1
	}

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	fmt.Printf("max = %d\n", max)
}
