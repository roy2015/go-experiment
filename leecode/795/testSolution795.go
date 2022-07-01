package main

import "fmt"

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	var sum int
	var pre bool
	var cnt int
	var current bool
	for _, num := range nums {
		current = num <= right
		if !current {
			if pre {
				sum += cnt * (cnt + 1) / 2
				pre = false
				cnt = 0
			} else {
			}
		} else {
			cnt++
			pre = true
		}
	}
	if current {
		sum += cnt * (cnt + 1) / 2
	}
	return sum
}

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{1, 1, 2, 1, 4, 3, 1}, 2, 3)) //8
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))       //7
}
