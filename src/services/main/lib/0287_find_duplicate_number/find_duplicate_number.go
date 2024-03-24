package findduplicatenumber

func FindDuplicate(nums []int) int {
	slow := nums[nums[0]]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
