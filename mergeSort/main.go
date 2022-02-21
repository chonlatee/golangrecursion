package main

import "fmt"

func mergeSortRecursion(l []int, start, end int) { // recursion

	if start < end {
		mid := (start + end) / 2
		mergeSortRecursion(l, start, mid)
		mergeSortRecursion(l, mid+1, end)
		merge(l, start, mid, end)
	}

}

// ref: https://github.com/HSatwick/ProgrammersCorner/blob/master/Merge%20Sort/mergeSort.java
func mergeSort(l []int) { // loop
	i := 2
	size := len(l) * 2
	for i < size {
		j := 0
		for j < len(l)-1 {
			start := j
			end := j + i - 1

			if end > len(l) {
				end = len(l) - 1
			}

			mid := (start + end) / 2

			merge(l, start, mid, end)
			j = j + i
		}

		i = i * 2

		//if i >= len(l) { // uncomment this block when change size = len(i)
		//	i = i / 2
		//	merge(l, 0, i-1, len(l)-1)
		//	i = len(l)
		//}
	}
}

func merge(l []int, start, mid, end int) { // merge
	size := end - start + 1
	temp := make([]int, size)

	i := start
	j := mid + 1
	k := 0

	for i <= mid && j <= end {
		if l[i] <= l[j] {
			temp[k] = l[i]
			k++
			i++
		} else {
			temp[k] = l[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = l[i]
		k++
		i++
	}

	for j <= end {
		temp[k] = l[j]
		k++
		j++
	}

	for n := start; n <= end; n++ {
		l[n] = temp[n-start]
	}

}

func main() {

	l := []int{0, -6, 10, -7, 9, -15, 20, -20}
	mergeSortRecursion(l, 0, len(l)-1)
	fmt.Printf("l=%v\n", l)

	k := []int{0, -6, 10, -7, 9, -15, 20, -20}
	mergeSort(k)
	fmt.Printf("k=%v\n", k)

}
