package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func PointerFromInt(val int) *int {
	prt := val
	return &prt
}

func CheckIfALevelIsEmpty(level *[]*TreeNode) bool {
	for _, y := range *level {
		if y != nil {
			return false
		}
	}
	return true
}

func NumberIsEven(n *int) bool {
	return *n&1 == 0
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func GetLinkedListLength(head *ListNode) int {
	var tail *ListNode = nil
	tail = head
	it := 0
	for tail.Next != nil {
		tail = tail.Next
		it += 1
	}
	return it + 1
}

func GenerateLinkedList(vals *[]int) *ListNode {
	if len(*vals) == 0 {
		return nil
	}
	head := NewNode((*vals)[0])
	var curr *ListNode = nil
	i := 1
	curr = head
	for ; i < len(*vals); i += 1 {
		curr.Next = NewNode((*vals)[i])
		curr = curr.Next
	}

	return head
}
