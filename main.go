package main

import (
	"fmt"
	"sort"
)

func sortDesc(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println("Sorted desc numbers:", numbers)
}

func sortAsc(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("Sorted asc numbers:", numbers)
}

func findMax(numbers []int) int {
	if len(numbers) > 0 {
		var result = numbers[0]
		for i := 1; i < len(numbers); i++ {
			if numbers[i] > result {
				result = numbers[i]
			}
		}
		return result
	} else {
		return 0
	}
}

func findMin(numbers []int) int {
	if len(numbers) > 0 {
		var result = numbers[0]
		for i := 1; i < len(numbers); i++ {
			if numbers[i] < result {
				result = numbers[i]
			}
		}
		return result
	} else {
		return 0
	}
}

func avgCalc(numbers []int) float64 {
	if len(numbers) > 0 {
		var result = 0.0
		for i := 0; i < len(numbers); i++ {
			result += float64(numbers[i]) / float64(len(numbers))
		}
		return result
	} else {
		return 0
	}
}

func main() {
	var numbers []int
	var count int
	fmt.Println("This is simple Go program.")
	fmt.Println("Please enter count of numbers array:")
	_, err := fmt.Scan(&count)
	if err != nil {
		count = 5
		fmt.Println("Set default array value to 5.")
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Please input %d number:", i+1)
		numbers = append(numbers, 0)
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			numbers[i] = 0
			continue
		}
	}
	fmt.Println("Input numbers:", numbers)
	sortDesc(numbers)
	sortAsc(numbers)
	fmt.Printf("The max number is %d.\n", findMax(numbers))
	fmt.Printf("The min number is %d.\n", findMin(numbers))
	fmt.Printf("The avg number is %.2f.\n", avgCalc(numbers))
}
