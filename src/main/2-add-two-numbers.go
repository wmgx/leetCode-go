package main

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	printList(addTwoNumbers(buildListNode([]int{2, 4, 3}), buildListNode([]int{5, 6, 4})))
	printList(addTwoNumbers(buildListNode([]int{0}), buildListNode([]int{0})))
	printList(addTwoNumbers(buildListNode([]int{9, 9, 9, 9, 9, 9, 9}), buildListNode([]int{9, 9, 9, 9})))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	next := head
	j := 0

	for l1 != nil || l2 != nil || j != 0 {
		t := &ListNode{j, nil}
		next.Next = t
		if l1 != nil {
			t.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t.Val += l2.Val
			l2 = l2.Next
		}
		j = t.Val / 10
		t.Val %= 10
		next = next.Next
	}
	return head.Next
}
