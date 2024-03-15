package productofarrayexceptself

// [1,2,3,4,5]
// [120,60,20]
// [1,1,2,6,24]
// [120,60,40,30,24]
// [60(1*3*4*5),40(2*4*5),30(6*5),24(ultimo)]
//

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	zeros := VerifyZeros(&nums)
	switch zeros {
	case 0:
		HandleNoZerosProd(&nums, &result)
	case 1:
		HandleOneZero(&nums, &result)
	case 2:
		return result
	default:
		return result
	}
	return result
}

func HandleNoZerosProd(nums *[]int, result *[]int) {
	if len(*nums) == 2 {
		(*result)[0], (*result)[1] = (*nums)[1], (*nums)[0]
		return
	}
	// left side
	(*result)[0], (*result)[1] = (*nums)[0], (*nums)[0]
	i := 2
	for ; i < len(*nums); i += 1 {
		(*result)[i] = (*nums)[i-1] * (*result)[i-1]
	}
	i = len(*nums) - 3
	auxArr := make([]int, len(*nums)-2)
	// add last el
	auxArr[i] = (*nums)[i+1] * (*nums)[i+2]
	i -= 1
	for ; i >= 0; i -= 1 {
		auxArr[i] = auxArr[i+1] * (*nums)[i+1]
	}
	i = len(*nums) - 2
	(*result)[i] = (*result)[i] * (*nums)[i+1]
	// result already have last 2 correct elements
	(*result)[0] = auxArr[0]
	i -= 1
	// calculate the rest only by multiplying two numbers
	for ; i > 0; i -= 1 {
		(*result)[i] = (*result)[i] * auxArr[i]
	}
}

func VerifyZeros(nums *[]int) uint8 {
	zeros := uint8(0)
	for _, x := range *nums {
		if x == 0 {
			zeros += 1
			if zeros > 1 {
				break
			}
		}
	}
	return zeros
}

func HandleOneZero(nums *[]int, result *[]int) {
	prod, i, x, zeroIndex := 1, 0, 0, 0
	for i, x = range *nums {
		if x == 0 {
			zeroIndex = i
			break
		}
	}
	i = 0
	for ; i < zeroIndex; i += 1 {
		prod *= (*nums)[i]
	}
	i = zeroIndex + 1
	for ; i < len(*nums); i += 1 {
		prod *= (*nums)[i]
	}
	(*result)[zeroIndex] = prod
}
