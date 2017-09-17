package main

func main() {
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	node1 := ListNode{2, nil}
	node2 := ListNode{4, nil}
	node3 := ListNode{3, nil}

	node1.Next = &node2
	node2.Next = &node3

	nodeA := ListNode{5, nil}
	nodeB := ListNode{6,nil}
	nodeC := ListNode{4, nil}

	nodeA.Next = &nodeB
	nodeB.Next = &nodeC

	addTwoNumbers(&node1, &nodeA)
}



