package lib

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	result := true
	stack1, stack2 := []*TreeNode{}, []*TreeNode{}
	curr1 := p
	curr2 := q
	var popedItem1 *TreeNode = nil
	var popedItem2 *TreeNode = nil
	shouldStop := false
	shouldStop = VerifyIfCurrNodeAreDifferents(curr1, curr2)
	if shouldStop {
		return false
	}
	for (curr1 != nil || len(stack1) > 0) &&
		(curr2 != nil || len(stack2) > 0) {

		for curr1 != nil || curr2 != nil {
			stack1 = append(stack1, curr1)
			curr1 = curr1.Left
			stack2 = append(stack2, curr2)
			curr2 = curr2.Left

			shouldStop = VerifyIfCurrNodeAreDifferents(curr1, curr2)
			if shouldStop {
				return false
			}
		}
		popedItem1 = stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		curr1 = popedItem1.Right

		popedItem2 = stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		curr2 = popedItem2.Right

		shouldStop = VerifyIfCurrNodeAreDifferents(curr1, curr2)
		if shouldStop {
			return false
		}
	}
	if len(stack1) != len(stack2) {
		result = false
	}
	return result
}

func VerifyIfCurrNodeAreDifferents(curr1 *TreeNode, curr2 *TreeNode) bool {
	if (curr1 == nil && curr2 != nil) ||
		(curr1 != nil && curr2 == nil) {
		return true
	}
	if curr1 != nil && curr2 != nil {
		if curr1.Val != curr2.Val {
			return true
		}
	}
	return false
}

func NewBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return &TreeNode{Val: arr[0], Left: nil, Right: nil}
}

func (t *TreeNode) Insert(val int) *TreeNode {
	var current, parent *TreeNode = nil, nil
	var node TreeNode = TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	current = t
	for {
		parent = current
		if val < current.Val {
			current = current.Left
			if current == nil {
				parent.Left = &node
				return t
			}
		} else {
			current = current.Right
			if current == nil {
				parent.Right = &node
				return t
			}
		}
	}
}
