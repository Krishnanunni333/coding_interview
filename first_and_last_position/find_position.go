package main

import "fmt"

func findFirstAndLast(arr []int, target int) [2]int {
	pos := [2]int{-1, -1}
	if arr[0] == target {
		return [2]int{0, 0}
	}
	var left int = 0
	var right int = len(arr) - 1
	var firstPos int = -1
	for left <= right {
		m := left + (right-left)/2
		fmt.Println(m)
		if arr[m] < target {
			left = m + 1
		} else if arr[m] > target {
			right = m - 1
		} else {
			if m == 0 || arr[m-1] != target {
				firstPos = m
				break
			} else {
				right = m - 1
			}
		}
	}
	if arr[len(arr) - 1] == target {
		return [2]int{firstPos, len(arr) - 1}
	}
	left = 0
	right = len(arr) - 1
	var secondPos int = -1
	
	for left <= right {
		m := left + (right-left)/2
		if target > arr[m] {
			left = m + 1
		} else if target < arr[m] {
			right = m - 1
		}else{
			if m == len(arr) - 1  || arr[m + 1] != target{
				secondPos = m
				break
			}else {
				left = m + 1
			}
		}
	}
	pos[0] = firstPos
	pos[1] = secondPos
	return pos
}
func main() {
	arr := []int{2, 4, 5, 5, 5, 5, 5, 7, 9, 9}
	fmt.Println(findFirstAndLast(arr, 9))
}
