package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
