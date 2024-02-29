package evenoddtree

import (
	"project/src/services/main/lib/utils"
)

type TreeNode = utils.TreeNode

func IsEvenOddTree(root *TreeNode) bool {

	if root == nil {
		return false
	}
	if NumberIsEven(&root.Val) {
		return false
	}
	isEvenLevel := false
	intLevel := 0
	treeLevels := GetTreeLevels(root)
	currLevelIter, currLevelLength, level := 0, 0, uint32(1)
	var currNode *TreeNode = nil
	treeHeight := uint32(len(*treeLevels) - 1)
	for {
		if level >= treeHeight {
			break
		}
		currLevelLength = len(*(*treeLevels)[level])
		// fmt.Printf("Level %d -> \n", level)
		intLevel = int(level)
		isEvenLevel = NumberIsEven(&intLevel)
		lastNonNilNodeIndex, previousValue := 0, 0
		for ; currLevelIter < currLevelLength; currLevelIter += 1 {
			currNode = (*(*treeLevels)[level])[currLevelIter]
			if currNode != nil {
				if isEvenLevel {
					if NumberIsEven(&currNode.Val) {
						return false
					}
					lastNonNilNodeIndex = GetLastNonNillPreviousNodeCompare((*treeLevels)[level], currLevelIter)
					if lastNonNilNodeIndex != -1 {
						previousValue = (*(*treeLevels)[level])[lastNonNilNodeIndex].Val
						if previousValue >= currNode.Val {
							return false
						}
					}

				} else {
					if !NumberIsEven(&currNode.Val) {
						return false
					}
					lastNonNilNodeIndex = GetLastNonNillPreviousNodeCompare((*treeLevels)[level], currLevelIter)
					if lastNonNilNodeIndex != -1 {
						previousValue = (*(*treeLevels)[level])[lastNonNilNodeIndex].Val
						if previousValue <= currNode.Val {
							return false
						}
					}
				}
			}
		}
		currLevelIter = 0
		level += 1
	}

	return true
}

// se retornar -1 o currNode é o primeiro, logo
// não necessita validar com o anterior, ou existe apenas ele no level
// qq retorno diferente de -1, é o indice do ultimo nó não nulo anterior
func GetLastNonNillPreviousNodeCompare(level *[]*TreeNode, currNodeIndex int) int {
	if currNodeIndex == 0 {
		return -1
	}
	var node *TreeNode = nil
	it := currNodeIndex
	for {
		it -= 1
		node = (*level)[it]
		if node != nil {
			return it
		}
		if it == 0 {
			break
		}
	}
	return -1
}

func NumberIsEven(n *int) bool {
	return *n&1 == 0
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Right: nil,
		Left:  nil,
	}
}

func GetTreeLevels(root *TreeNode) *map[uint32]*[]*TreeNode {
	treeLevels := map[uint32]*[]*TreeNode{}
	treeLevels[0] = &[]*TreeNode{root}
	if root.Left == nil && root.Right == nil {
		return &treeLevels
	}
	level := 1
	var currLevel *[]*TreeNode = nil
	for {
		treeLevels[uint32(level)] = &[]*TreeNode{}
		for _, x := range *treeLevels[uint32(level-1)] {
			if x != nil {
				*treeLevels[uint32(level)] = append(*treeLevels[uint32(level)], x.Left)
				*treeLevels[uint32(level)] = append(*treeLevels[uint32(level)], x.Right)
			}
		}
		currLevel = treeLevels[uint32(level)]
		if utils.CheckIfALevelIsEmpty(currLevel) {
			break
		}
		level += 1
	}
	return &treeLevels
}
