package main

import (
	"container/list"
	"fmt"
)

func reverseStringRecursion(s string) string {
	if s == "" {
		return s
	}

	return reverseStringRecursion(s[1:]) + string(s[0])
}

func reverseStringStack(s string) string {
	st := list.New()
	ss := ""
	for _, v := range s {
		st.PushBack(string(v))
	}

	for st.Len() != 0 {
		e := st.Remove(st.Back())
		ss += e.(string)
	}

	return ss

}

func reverseString(s string) string {
	ss := ""
	for i := len(s) - 1; i >= 0; i-- {
		ss += string(s[i])
	}

	return ss
}

func main() {

	s := "abcdef"

	fmt.Println("recusion: ", reverseStringRecursion(s))

	fmt.Println("stack: ", reverseStringStack(s))

	fmt.Println("for loop: ", reverseString(s))
}
