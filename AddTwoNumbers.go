package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	headNode1 := l1
	headNode2 := l2
	sum := l1.Val+l2.Val
	if sum >= 10{
		add = 1
	} else {
		add = 0
	}
	headNode := ListNode{sum%10, nil}
	currentNode := &headNode

	for {

		if (headNode1.Next == nil && headNode2.Next == nil)&& add == 0 {
			break
		} else {

			if headNode1.Next == nil {
				headNode1 = &ListNode{0, nil}
			} else {
				headNode1 = headNode1.Next
			}

			if headNode2.Next == nil {
				headNode2 = &ListNode{0, nil}
			} else {
				headNode2 = headNode2.Next
			}

		}
			sum = headNode1.Val + headNode2.Val + add
			add = 0

			newNode := ListNode{sum%10, nil}

			if sum >= 10{
				add = 1
			} else {
				add = 0
			}

			tempNode := currentNode
			currentNode = &newNode
			tempNode.Next = currentNode

	}

	return &headNode

}
