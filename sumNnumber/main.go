package main

import (
	"container/list"
	"fmt"
)

func sumRecursion(input int) int { // use recursion
	if input <= 1 {
		return input
	}

	return input + sumRecursion(input-1)
}

func sum(input int) int { // use loop
	n := 0
	for i := 1; i <= input; i++ {
		n += i
	}

	return n
}

func sumStack(input int) int { // use stack
	n := 0
	st := list.New()
	for input >= 1 {
		st.PushBack(input)
		input = input - 1
	}

	for st.Len() > 0 {
		e := st.Back()
		v := st.Remove(e)

		n += v.(int)
	}

	return n
}

func sumGauss(input int) int {
	return (input * (input + 1)) / 2 // gauss formula
}

func main() {

	num := 10

	fmt.Println(sumRecursion(num))
	fmt.Println(sum(num))
	fmt.Println(sumStack(num))
	fmt.Println(sumGauss(num))

}
