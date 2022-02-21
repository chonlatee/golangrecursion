package main

import (
	"container/list"
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func binarySearchTreeRecursion(head *node, val int) *node { // recursion
	if head == nil {
		n := &node{
			val: val,
		}
		return n
	}

	if head.val < val {
		head.right = binarySearchTreeRecursion(head.right, val)
	} else {
		head.left = binarySearchTreeRecursion(head.left, val)
	}

	return head
}

func binarySearchTree(head *node, val int) *node { // loop
	newNode := &node{
		val: val,
	}
	n := head
	for {
		if n.val < val {
			if n.right == nil {
				n.right = newNode
				return head
			}
			n = n.right
		} else {
			if n.left == nil {
				n.left = newNode
				return head
			}
			n = n.left
		}
	}
}

func binarySearchTreeII(head *node, val int) *node { // loop but safer

	newNode := &node{
		val: val,
	}
	n := head
	var prev *node
	for n != nil {
		prev = n
		if n.val < val {
			n = n.right
		} else {
			n = n.left
		}
	}

	if prev.val < val {
		prev.right = newNode
	} else {
		prev.left = newNode
	}

	return head
}

func printLeavesRecursion(root *node) { // recursion
	if root == nil {
		return
	}

	if root.isLeaf() {
		fmt.Printf("[%v] ", root.val)
		return
	}

	if root.left != nil {
		printLeavesRecursion(root.left)
	}

	if root.right != nil {
		printLeavesRecursion(root.right)
	}
}

func printLeavesQueue(root *node) { // queue

	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		v := st.Front().Value
		n := v.(*node)

		if n.left != nil {
			st.PushBack(n.left)
		}

		if n.right != nil {
			st.PushBack(n.right)
		}

		if n.left == nil && n.right == nil {
			fmt.Printf("[%v] ", n.val)
		}

		st.Remove(st.Front())
	}
}

func printLeavesStackBFS(root *node) { // stack
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		v := st.Back().Value
		st.Remove(st.Back())
		n := v.(*node)

		if n.right != nil {
			st.PushBack(n.right)
		}

		if n.left != nil {
			st.PushBack(n.left)
		}

		if n.left == nil && n.right == nil {
			fmt.Printf("[%v] ", n.val)
		}
	}
}

func printLeavesStackDFS(root *node) { // stack dfs
	st := list.New()
	st.PushBack(root)
	visited := make(map[*node]bool)
	visited[root] = true
	for st.Len() > 0 { // stack is empty?
		v := st.Back().Value
		n := v.(*node)

		if n.left != nil && !visited[n.left] {
			st.PushBack(n.left)
			visited[n.left] = true
			continue
		}

		if n.right != nil && !visited[n.right] {
			st.PushBack(n.right)
			visited[n.right] = true
			continue
		}

		if n.left == nil && n.right == nil {
			fmt.Printf("[%v] ", n.val)
		}

		st.Remove(st.Back())
	}
}

func main() {

	root := &node{
		val:   100,
		left:  nil,
		right: nil,
	}

	h1 := binarySearchTreeRecursion(root, 80)
	binarySearchTreeRecursion(root, 120)
	binarySearchTreeRecursion(root, 50)
	binarySearchTreeRecursion(root, 90)
	binarySearchTreeRecursion(root, 110)
	binarySearchTreeRecursion(root, 140)
	binarySearchTreeRecursion(root, 30)
	binarySearchTreeRecursion(root, 60)
	binarySearchTreeRecursion(root, 85)
	binarySearchTreeRecursion(root, 95)
	binarySearchTreeRecursion(root, 108)
	binarySearchTreeRecursion(root, 115)
	binarySearchTreeRecursion(root, 150)

	printLeavesRecursion(h1)
	fmt.Println()

	root = &node{
		val: 100,
	}

	h2 := binarySearchTree(root, 80)
	binarySearchTree(root, 120)
	binarySearchTree(root, 50)
	binarySearchTree(root, 90)
	binarySearchTree(root, 110)
	binarySearchTree(root, 140)
	binarySearchTree(root, 30)
	binarySearchTree(root, 60)
	binarySearchTree(root, 85)
	binarySearchTree(root, 95)
	binarySearchTree(root, 108)
	binarySearchTree(root, 115)
	binarySearchTree(root, 150)

	printLeavesQueue(h2)
	fmt.Println()

	root = &node{
		val: 100,
	}

	h3 := binarySearchTreeII(root, 80)
	binarySearchTreeII(root, 120)
	binarySearchTreeII(root, 50)
	binarySearchTreeII(root, 90)
	binarySearchTreeII(root, 110)
	binarySearchTreeII(root, 140)
	binarySearchTreeII(root, 30)
	binarySearchTreeII(root, 60)
	binarySearchTreeII(root, 85)
	binarySearchTreeII(root, 95)
	binarySearchTreeII(root, 108)
	binarySearchTreeII(root, 115)
	binarySearchTreeII(root, 150)

	printLeavesStackBFS(h3)
	fmt.Println()

	root = &node{
		val: 100,
	}

	h4 := binarySearchTree(root, 80)
	binarySearchTree(root, 120)
	binarySearchTree(root, 50)
	binarySearchTree(root, 90)
	binarySearchTree(root, 110)
	binarySearchTree(root, 140)
	binarySearchTree(root, 30)
	binarySearchTree(root, 60)
	binarySearchTree(root, 85)
	binarySearchTree(root, 95)
	binarySearchTree(root, 108)
	binarySearchTree(root, 115)
	binarySearchTree(root, 150)

	printLeavesStackDFS(h4)
	fmt.Println()
}
