package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	h := 0
	return DiameterOptimized(root, &h)
}

func DiameterOptimized(root *TreeNode, h *int) int {
	lh, rh := 0, 0
	lDiameter, rDiameter := 0, 0

	if root == nil {
		*h = 0
		return 0
	}

	lDiameter = DiameterOptimized(root.Left, &lh)
	rDiameter = DiameterOptimized(root.Right, &rh)
	*h = Max(lh, rh) + 1
	return Max(lh+rh+1, Max(lDiameter, rDiameter)) - 1
}

func TreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(TreeHeight(root.Left), TreeHeight(root.Right)) + 1
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return &TreeNode{Val: arr[0], Left: nil, Right: nil}
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Right: nil,
		Left:  nil,
	}
}

func (t *TreeNode) Insert(val int) *TreeNode {
	var current, parent *TreeNode = nil, nil
	node := TreeNode{
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
