package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func decimalToBinaryRecursion(input int, result string) string {

	if input == 0 {
		return result
	}

	result = strconv.Itoa(input%2) + result

	return decimalToBinaryRecursion(input/2, result)
}

func decimalToBinary(input int) string {
	res := ""
	for input > 0 {
		res = strconv.Itoa(input%2) + res
		input = input / 2
	}
	return res
}

func decimalToBinaryStack(input int) string {
	res := ""

	stack := list.New()

	for input > 0 {
		stack.PushBack(strconv.Itoa(input % 2)) // push
		input = input / 2
	}

	for stack.Len() > 0 {
		v := stack.Back()
		e := stack.Remove(v) // pop
		res = res + e.(string)
	}

	return res
}

func decimalToBinaryQueue(input int) string {
	res := ""

	queue := list.New()

	for input > 0 {
		queue.PushBack(strconv.Itoa(input % 2)) // enqueue
		input = input / 2
	}

	for queue.Len() > 0 {
		v := queue.Front()
		e := queue.Remove(v) // dequeue
		res = e.(string) + res
	}

	return res
}

func main() {

	num := 18
	result := ""

	res := decimalToBinaryRecursion(num, result)

	fmt.Println(res)

	fmt.Println(decimalToBinary(num))

	fmt.Println(decimalToBinaryStack(num))

	fmt.Println(decimalToBinaryQueue(num))

}
