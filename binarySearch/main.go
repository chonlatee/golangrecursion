package main

import "fmt"

func binarySearchRecursion(l []int, left int, right int, x int) int { // use recursion
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if x == l[mid] {
		return mid
	}

	if x < l[mid] {
		return binarySearchRecursion(l, left, mid-1, x)
	}

	return binarySearchRecursion(l, mid+1, right, x)
}

func binarySearch(l []int, x int) int { // use loop
	left := 0
	right := len(l) - 1

	for left < right {
		mid := (left + right) / 2
		if x == l[mid] {
			return mid
		}

		if x > l[mid] {
			left = mid + 1
		}

		if x < l[mid] {
			right = mid - 1
		}
	}

	return -1
}

func main() {

	l := []int{1, 2, 4, 6, 8, 9, 20, 25, 30, 48, 50}

	i := binarySearchRecursion(l, 0, len(l)-1, 20)
	fmt.Println(i)

	j := binarySearch(l, 20)
	fmt.Println(j)

}
