package mergeinbetweenlinkedlists

import "project/src/services/main/lib/utils"

type ListNode = utils.ListNode

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	nodeBeforeA := list1
	i := 0
	for ; i < a-1; i += 1 {
		nodeBeforeA = nodeBeforeA.Next
	}
	nodeB := nodeBeforeA.Next
	i = 0
	for ; i < b-a; i += 1 {
		nodeB = nodeB.Next
	}
	nodeBeforeA.Next = list2
	lastInlist2 := list2
	for lastInlist2.Next != nil {
		lastInlist2 = lastInlist2.Next
	}
	lastInlist2.Next = nodeB.Next
	nodeB.Next = nil

	return list1
}
