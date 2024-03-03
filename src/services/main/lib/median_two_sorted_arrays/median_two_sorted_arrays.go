package mediantwosortedarrays

import (
	"project/src/services/main/lib/utils"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := float64(0)
	size1, size2 := len(nums1), len(nums2)
	totalMergedLen := size1 + size2
	if totalMergedLen == 0 {
		return result
	}
	posMedianMerged := 0
	mergedSlice := []int{}
	mergedIsEven := utils.NumberIsEven(&totalMergedLen)
	if mergedIsEven {
		posMedianMerged = totalMergedLen / 2
	} else {
		posMedianMerged = (totalMergedLen - 1) / 2
	}
	nums1Iter, nums2Iter := 0, 0
	hasRoomForward1, hasRoomForward2 := false, false
	maxLimit := posMedianMerged + 1
	for {
		hasRoomForward1 = nums1Iter < size1
		hasRoomForward2 = nums2Iter < size2
		if len(mergedSlice) >= maxLimit || !hasRoomForward1 && !hasRoomForward2 {
			break
		}
		if hasRoomForward1 && hasRoomForward2 {
			if nums1[nums1Iter] < nums2[nums2Iter] {
				mergedSlice = append(mergedSlice, nums1[nums1Iter])
				nums1Iter += 1
			} else if nums1[nums1Iter] > nums2[nums2Iter] {
				mergedSlice = append(mergedSlice, nums2[nums2Iter])
				nums2Iter += 1
			} else {
				if len(mergedSlice) < maxLimit {
					mergedSlice = append(mergedSlice, nums2[nums2Iter])
					nums2Iter += 1
				} else {
					break
				}
				if len(mergedSlice) < maxLimit {
					mergedSlice = append(mergedSlice, nums1[nums1Iter])
					nums1Iter += 1
				} else {
					break
				}

			}
		} else if hasRoomForward1 && !hasRoomForward2 {
			mergedSlice = append(mergedSlice, nums1[nums1Iter])
			nums1Iter += 1
		} else if hasRoomForward2 && !hasRoomForward1 {
			mergedSlice = append(mergedSlice, nums2[nums2Iter])
			nums2Iter += 1
		}
	}
	lenMerged := len(mergedSlice)
	if lenMerged == 1 {
		result = float64(mergedSlice[0])
		return result
	}

	if mergedIsEven {
		sum := float64(mergedSlice[lenMerged-1]) + float64(mergedSlice[lenMerged-2])
		if sum != 0 {
			result = sum / 2
		} else {
			result = 0
		}
	} else {
		result = float64(mergedSlice[lenMerged-1])
	}

	return result
}
