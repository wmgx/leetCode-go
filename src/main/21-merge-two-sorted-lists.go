package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

func main() {
	printList(mergeTwoLists(buildListNode([]int{1, 2, 4}), buildListNode([]int{1, 3, 4})))
	printList(mergeTwoLists(buildListNode([]int{}), buildListNode([]int{})))
	printList(mergeTwoLists(buildListNode([]int{}), buildListNode([]int{0})))
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = new(ListNode)
	var next = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next = list2
			list2 = list2.Next
		}
		next = next.Next
	}
	var last *ListNode
	if list1 == nil {
		last = list2
	} else {
		last = list1
	}
	if last != nil {
		next.Next = last
	}
	return head.Next
}

func IF(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
