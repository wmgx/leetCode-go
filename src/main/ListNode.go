package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(nums []int) *ListNode {
	var head = new(ListNode)
	var next = head
	for _, v := range nums {
		next.Next = &ListNode{v, nil}
		next = next.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	var next = head
	if next == nil {
		fmt.Println("empty List")
		return
	}
	for next != nil {
		fmt.Print(next.Val, " ")
		next = next.Next
	}
	fmt.Println()
}
func reverse(list *ListNode) *ListNode {
	var head = new(ListNode)

	for list != nil {
		var next = list.Next
		list.Next = head.Next
		head.Next = list
		list = next
	}
	return head.Next
}
