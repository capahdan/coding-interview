package product_except_self

func ProductExceptSelf(nums []int) []int {

	length := len(nums)
	result := make([]int, length)

	// Left products
	leftProduct := 1
	for i := 0; i < length; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Right products
	rightProduct := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

//
// solusi kedua yang muncul di otak ku adalah
// kita bikin ajah sebuah for loop yang akan mengisikan semua index dari nums dan terakhir menghapus indexnya sendiri

// map[0][1,2,3]
// map[1][0,2,3]
// map[2][0,1,3]
// map[3][0,1,2]
