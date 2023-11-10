package search_insert

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

// untuk hadle positive flownya
// masukan semua array kedalam map ,
// kalo di java script yang tinggal num.indexOf(target)
// yang kalo ada dia bakal return langsung tapi kalo gak ada ya di bakalan return -1
//
