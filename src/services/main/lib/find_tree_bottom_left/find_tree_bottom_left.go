package findtreebottomleft

import (
	"project/src/services/main/lib/utils"
)

type TreeNode = utils.TreeNode

func FindBottomLeftValue(root *TreeNode) int {
	lastLevel := GetTreeLeafs(root)
	i := 0
	for {
		if lastLevel[i] != nil {
			i = lastLevel[i].Val
			break
		}
	}
	return i
}

func GetTreeLeafs(root *TreeNode) []*TreeNode {
	treeLevels := map[uint8]*[]*TreeNode{}
	treeLevels[0] = &[]*TreeNode{root}
	if root.Left == nil && root.Right == nil {
		return *treeLevels[0]
	}
	lastLeafLevel, level := 0, 1

	var currLevel *[]*TreeNode = nil
	for {
		treeLevels[uint8(level)] = &[]*TreeNode{}
		for _, x := range *treeLevels[uint8(level-1)] {
			if x != nil {
				*treeLevels[uint8(level)] = append(*treeLevels[uint8(level)], x.Left)
				*treeLevels[uint8(level)] = append(*treeLevels[uint8(level)], x.Right)
			}
		}
		currLevel = treeLevels[uint8(level)]
		if utils.CheckIfALevelIsEmpty(currLevel) {
			break
		}
		lastLeafLevel = level
		level += 1
	}
	return *treeLevels[uint8(lastLeafLevel)]
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Right: nil,
		Left:  nil,
	}
}
