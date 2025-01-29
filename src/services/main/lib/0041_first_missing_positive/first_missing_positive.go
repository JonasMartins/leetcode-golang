package firstmissingpositive

func Firstmissingpositive(nums []int) int {
	return nums[0]
}

func Swap(arr *[]int, pos1 int, pos2 int) {
	(*arr)[pos1] = (*arr)[pos2]
}
