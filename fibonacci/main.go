package main

import "fmt"

var cache = make(map[int]int)

func fibRecursion(n int) int { // recursion
	if n == 0 || n == 1 {
		return n
	}

	return fibRecursion(n-1) + fibRecursion(n-2)
}

func fib(n int) int {
	i := 0
	a := 0
	b := 1
	for i < n {
		temp := a
		a = b
		b = temp + b // short: a, b = b, a + b
		i++
	}

	return a
}

func fibRecursionMemo(n int) int { // memo
	if n == 0 || n == 1 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	} else {
		r := fibRecursionMemo(n-1) + fibRecursionMemo(n-2)
		cache[n] = r
		return r
	}
}

func main() {
	num := 50

	fmt.Println(fib(num))
	fmt.Println(fibRecursion(num))
	fmt.Println(fibRecursionMemo(num))
}
