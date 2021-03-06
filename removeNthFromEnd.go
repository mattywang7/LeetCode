package main

import "fmt"

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Input:	head = [1,2,3,4,5], n = 2
// Output:	[1, 2, 3, 5]

// Input:	head = [1], n = 1
// Output:	[]

// Input:	head = [1, 2], n = 1
// Output:	[1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := createList(2)
	head = removeNthFromEnd(head, 1)
	printListNode(head)
}

// from head: Length - n + 1
// 0ms
// 2.2MB
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	tmp = &ListNode{0, head}
	for i := 0; i < length - n; i++ {
		tmp = tmp.Next
	}
	if tmp.Next == head {
		head = head.Next
	} else {
		tmp.Next = tmp.Next.Next
	}
	return head
}

func createList(n int) *ListNode {
	head := &ListNode{1, nil}
	tmp := head
	for i := 2; i <= n; i++ {
		tmp.Next = &ListNode{i, nil}
		tmp = tmp.Next
	}
	return head
}

func printListNode(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

// One pass
// 0ms
// 2.2MB
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy
	for i := 1; i <= n + 1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
